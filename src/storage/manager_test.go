package storage

import (
	"testing"
)

func TestNewPage(t *testing.T) {
	page := NewPage(1)
	if page.ID != 1 {
		t.Errorf("Expected page ID 1, got %d", page.ID)
	}
}

func TestStorageManager(t *testing.T) {
	sm := NewStorageManager()
	
	// Allocate a page
	page := sm.AllocatePage()
	if page == nil {
		t.Fatal("Expected allocated page, got nil")
	}
	
	// Retrieve the page
	retrieved, exists := sm.GetPage(page.ID)
	if !exists {
		t.Error("Expected page to exist")
	}
	if retrieved.ID != page.ID {
		t.Errorf("Expected page ID %d, got %d", page.ID, retrieved.ID)
	}
}
