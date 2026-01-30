package transaction

import (
	"testing"
)

func TestTransactionManager(t *testing.T) {
	tm := NewTransactionManager()
	
	// Begin a transaction
	txn := tm.Begin()
	if txn.State != Active {
		t.Errorf("Expected transaction state Active, got %v", txn.State)
	}
	
	// Commit the transaction
	err := tm.Commit(txn)
	if err != nil {
		t.Errorf("Failed to commit transaction: %v", err)
	}
	if txn.State != Committed {
		t.Errorf("Expected transaction state Committed, got %v", txn.State)
	}
}

func TestWAL(t *testing.T) {
	wal := NewWAL()
	
	record := LogRecord{
		TransactionID: 1,
		Type:          "INSERT",
		Data:          []byte("test data"),
	}
	
	err := wal.Append(record)
	if err != nil {
		t.Errorf("Failed to append log record: %v", err)
	}
}
