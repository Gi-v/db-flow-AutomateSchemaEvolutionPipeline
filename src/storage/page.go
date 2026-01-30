package storage

import (
	"fmt"
)

// PageSize is the size of each page in bytes (8KB)
const PageSize = 8192

// PageID uniquely identifies a page
type PageID uint32

// Page represents a disk page in the buffer pool
type Page struct {
	ID       PageID
	Data     []byte
	Dirty    bool
	PinCount int
}

// NewPage creates a new page with the given ID
func NewPage(id PageID) *Page {
	return &Page{
		ID:       id,
		Data:     make([]byte, PageSize),
		Dirty:    false,
		PinCount: 0,
	}
}

// Pin increments the pin count for this page
func (p *Page) Pin() {
	p.PinCount++
}

// Unpin decrements the pin count for this page
func (p *Page) Unpin() error {
	if p.PinCount <= 0 {
		return fmt.Errorf("page %d already unpinned", p.ID)
	}
	p.PinCount--
	return nil
}

// IsPinned returns true if the page is pinned
func (p *Page) IsPinned() bool {
	return p.PinCount > 0
}

// MarkDirty marks the page as modified
func (p *Page) MarkDirty() {
	p.Dirty = true
}

// IsDirty returns true if the page has been modified
func (p *Page) IsDirty() bool {
	return p.Dirty
}

// Reset resets the page to initial state
func (p *Page) Reset() {
	p.Data = make([]byte, PageSize)
	p.Dirty = false
	p.PinCount = 0
}

// Write writes data to the page at the specified offset
func (p *Page) Write(offset int, data []byte) error {
	if offset < 0 || offset+len(data) > PageSize {
		return fmt.Errorf("write exceeds page size")
	}
	copy(p.Data[offset:], data)
	p.MarkDirty()
	return nil
}

// Read reads data from the page at the specified offset
func (p *Page) Read(offset int, length int) ([]byte, error) {
	if offset < 0 || offset+length > PageSize {
		return nil, fmt.Errorf("read exceeds page size")
	}
	result := make([]byte, length)
	copy(result, p.Data[offset:offset+length])
	return result, nil
}
