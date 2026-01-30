package storage

import (
	"fmt"
	"sync"
)

// BufferPool manages a pool of pages in memory
type BufferPool struct {
	pages    map[PageID]*Page
	capacity int
	mu       sync.RWMutex
	lru      []PageID // Simple LRU eviction policy
}

// NewBufferPool creates a new buffer pool with the specified capacity
func NewBufferPool(capacity int) *BufferPool {
	return &BufferPool{
		pages:    make(map[PageID]*Page),
		capacity: capacity,
		lru:      make([]PageID, 0, capacity),
	}
}

// FetchPage fetches a page from the buffer pool or disk
func (bp *BufferPool) FetchPage(pageID PageID) (*Page, error) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	// Check if page is already in buffer pool
	if page, exists := bp.pages[pageID]; exists {
		page.Pin()
		bp.updateLRU(pageID)
		return page, nil
	}

	// Evict a page if buffer pool is full
	if len(bp.pages) >= bp.capacity {
		if err := bp.evictPage(); err != nil {
			return nil, err
		}
	}

	// Load page from disk (stub: create new page)
	page := bp.loadPageFromDisk(pageID)
	page.Pin()
	bp.pages[pageID] = page
	bp.lru = append(bp.lru, pageID)

	return page, nil
}

// UnpinPage unpins a page in the buffer pool
func (bp *BufferPool) UnpinPage(pageID PageID, dirty bool) error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	page, exists := bp.pages[pageID]
	if !exists {
		return fmt.Errorf("page %d not in buffer pool", pageID)
	}

	if dirty {
		page.MarkDirty()
	}

	return page.Unpin()
}

// FlushPage writes a page to disk if it's dirty
func (bp *BufferPool) FlushPage(pageID PageID) error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	page, exists := bp.pages[pageID]
	if !exists {
		return fmt.Errorf("page %d not in buffer pool", pageID)
	}

	if page.IsDirty() {
		// Write page to disk (stub)
		bp.writePageToDisk(page)
		page.Dirty = false
	}

	return nil
}

// FlushAllPages flushes all dirty pages to disk
func (bp *BufferPool) FlushAllPages() error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	for _, page := range bp.pages {
		if page.IsDirty() {
			bp.writePageToDisk(page)
			page.Dirty = false
		}
	}

	return nil
}

// Size returns the current number of pages in the buffer pool
func (bp *BufferPool) Size() int {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return len(bp.pages)
}

// evictPage evicts a page using LRU policy
func (bp *BufferPool) evictPage() error {
	// Find first unpinned page in LRU list
	for i, pageID := range bp.lru {
		page := bp.pages[pageID]
		if !page.IsPinned() {
			// Flush if dirty
			if page.IsDirty() {
				bp.writePageToDisk(page)
			}
			// Remove from pool
			delete(bp.pages, pageID)
			bp.lru = append(bp.lru[:i], bp.lru[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("all pages are pinned, cannot evict")
}

// updateLRU updates the LRU list when a page is accessed
func (bp *BufferPool) updateLRU(pageID PageID) {
	// Move page to end of LRU list (most recently used)
	for i, id := range bp.lru {
		if id == pageID {
			bp.lru = append(bp.lru[:i], bp.lru[i+1:]...)
			break
		}
	}
	bp.lru = append(bp.lru, pageID)
}

// loadPageFromDisk loads a page from disk (stub implementation)
func (bp *BufferPool) loadPageFromDisk(pageID PageID) *Page {
	// TODO: Implement actual disk I/O
	return NewPage(pageID)
}

// writePageToDisk writes a page to disk (stub implementation)
func (bp *BufferPool) writePageToDisk(page *Page) {
	// TODO: Implement actual disk I/O
}
