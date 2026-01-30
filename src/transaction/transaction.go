package transaction

import (
	"sync"
)

// TransactionState represents the state of a transaction
type TransactionState int

const (
	Active TransactionState = iota
	Committed
	Aborted
)

// Transaction represents a database transaction
type Transaction struct {
	ID    int
	State TransactionState
	mu    sync.Mutex
}

// TransactionManager manages transactions
type TransactionManager struct {
	mu           sync.Mutex
	transactions map[int]*Transaction
	nextID       int
}

// NewTransactionManager creates a new transaction manager
func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		transactions: make(map[int]*Transaction),
		nextID:       1,
	}
}

// Begin starts a new transaction
func (tm *TransactionManager) Begin() *Transaction {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	
	txn := &Transaction{
		ID:    tm.nextID,
		State: Active,
	}
	tm.transactions[txn.ID] = txn
	tm.nextID++
	
	return txn
}

// Commit commits a transaction
func (tm *TransactionManager) Commit(txn *Transaction) error {
	txn.mu.Lock()
	defer txn.mu.Unlock()
	
	txn.State = Committed
	return nil
}

// Abort aborts a transaction
func (tm *TransactionManager) Abort(txn *Transaction) error {
	txn.mu.Lock()
	defer txn.mu.Unlock()
	
	txn.State = Aborted
	return nil
}
