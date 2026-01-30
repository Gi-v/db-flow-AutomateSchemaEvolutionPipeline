package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/executor"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/index"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/parser"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/storage"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/transaction"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Prometheus metrics
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "minidb_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)
	queryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "minidb_query_duration_seconds",
			Help:    "Query execution duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_type"},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(queryDuration)
}

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize components
	bufferPool := storage.NewBufferPool(100)
	tableManager := storage.NewTableManager(bufferPool)
	indexManager := index.NewBTreeManager()
	walManager := transaction.NewWALManager()
	lockManager := transaction.NewLockManager()
	sqlParser := parser.NewParser()
	queryExecutor := executor.NewExecutor(tableManager, indexManager, walManager, lockManager)

	// Log initialization
	log.Println("Mini Relational DBMS starting...")
	log.Printf("Buffer Pool: %d pages\n", bufferPool.Size())
	log.Printf("Components initialized: TableManager, IndexManager, WAL, LockManager, Parser, Executor")

	// Setup Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		requestsTotal.WithLabelValues("GET", "/health").Inc()
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"service": "minidb",
		})
	})

	// Metrics endpoint for Prometheus
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// SQL query endpoint
	router.POST("/query", func(c *gin.Context) {
		requestsTotal.WithLabelValues("POST", "/query").Inc()

		var request struct {
			SQL string `json:"sql" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse SQL
		ast, err := sqlParser.Parse(request.SQL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parse error: " + err.Error()})
			return
		}

		// Execute query
		timer := prometheus.NewTimer(queryDuration.WithLabelValues(ast.Type()))
		result, err := queryExecutor.Execute(ast)
		timer.ObserveDuration()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Execution error: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result": result,
		})
	})

	// API information endpoint
	router.GET("/api/info", func(c *gin.Context) {
		requestsTotal.WithLabelValues("GET", "/api/info").Inc()
		c.JSON(http.StatusOK, gin.H{
			"name":    "Mini Relational DBMS",
			"version": "0.1.0",
			"features": []string{
				"SQL: SELECT, INSERT, UPDATE, DELETE",
				"Storage: Disk-backed pages with buffer pool",
				"Indexing: B+ tree",
				"Transactions: WAL and MVCC",
				"Concurrency: Lock manager",
			},
		})
	})

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server starting on %s\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
