package storage

import (
	"sync"
)

// StorageManager manages page storage and retrieval
type StorageManager struct {
	mu    sync.RWMutex
	pages map[int]*Page
}

// NewStorageManager creates a new storage manager
func NewStorageManager() *StorageManager {
	return &StorageManager{
		pages: make(map[int]*Page),
	}
}

// GetPage retrieves a page by ID
func (sm *StorageManager) GetPage(id int) (*Page, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	page, exists := sm.pages[id]
	return page, exists
}

// AllocatePage allocates a new page
func (sm *StorageManager) AllocatePage() *Page {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	id := len(sm.pages)
	page := NewPage(id)
	sm.pages[id] = page
	
	return page
}
