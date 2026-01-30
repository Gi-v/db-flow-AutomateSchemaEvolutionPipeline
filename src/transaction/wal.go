package transaction

import (
	"sync"
	"time"
)

// WALEntry represents a Write-Ahead Log entry
type WALEntry struct {
	LSN       uint64      // Log Sequence Number
	Timestamp time.Time
	TxnID     uint64
	Operation string // "INSERT", "UPDATE", "DELETE", "COMMIT", "ABORT"
	TableName string
	Data      interface{}
}

// WALManager manages the Write-Ahead Log
type WALManager struct {
	entries []WALEntry
	nextLSN uint64
	nextTxn uint64
	mu      sync.Mutex
}

// NewWALManager creates a new WAL manager
func NewWALManager() *WALManager {
	return &WALManager{
		entries: make([]WALEntry, 0),
		nextLSN: 1,
		nextTxn: 1,
	}
}

// BeginTransaction starts a new transaction
func (w *WALManager) BeginTransaction() uint64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	txnID := w.nextTxn
	w.nextTxn++

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		TxnID:     txnID,
		Operation: "BEGIN",
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++

	return txnID
}

// LogInsert logs an INSERT operation
func (w *WALManager) LogInsert(tableName string, data interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		Operation: "INSERT",
		TableName: tableName,
		Data:      data,
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++
}

// LogUpdate logs an UPDATE operation
func (w *WALManager) LogUpdate(tableName string, setClause, where interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		Operation: "UPDATE",
		TableName: tableName,
		Data: map[string]interface{}{
			"set":   setClause,
			"where": where,
		},
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++
}

// LogDelete logs a DELETE operation
func (w *WALManager) LogDelete(tableName string, where interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		Operation: "DELETE",
		TableName: tableName,
		Data:      where,
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++
}

// Commit commits a transaction
func (w *WALManager) Commit(txnID uint64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		TxnID:     txnID,
		Operation: "COMMIT",
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++

	// TODO: Flush WAL to disk
	return nil
}

// Abort aborts a transaction
func (w *WALManager) Abort(txnID uint64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		LSN:       w.nextLSN,
		Timestamp: time.Now(),
		TxnID:     txnID,
		Operation: "ABORT",
	}
	w.entries = append(w.entries, entry)
	w.nextLSN++

	// TODO: Rollback changes
	return nil
}

// Recover replays the WAL for recovery
func (w *WALManager) Recover() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement ARIES-style recovery
	// 1. Analysis phase
	// 2. Redo phase
	// 3. Undo phase

	return nil
}

// Checkpoint creates a checkpoint
func (w *WALManager) Checkpoint() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement checkpointing
	// 1. Flush all dirty pages
	// 2. Write checkpoint record to WAL
	// 3. Truncate old WAL entries

	return nil
}

// GetEntries returns all WAL entries (for testing/debugging)
func (w *WALManager) GetEntries() []WALEntry {
	w.mu.Lock()
	defer w.mu.Unlock()

	entries := make([]WALEntry, len(w.entries))
	copy(entries, w.entries)
	return entries
}
