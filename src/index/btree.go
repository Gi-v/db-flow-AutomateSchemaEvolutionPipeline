package index

import (
	"fmt"
	"sync"
)

// BTreeOrder is the order of the B+ tree
const BTreeOrder = 4

// BTreeNode represents a node in the B+ tree
type BTreeNode struct {
	IsLeaf   bool
	Keys     []int
	Values   []interface{} // For leaf nodes
	Children []*BTreeNode  // For internal nodes
	Next     *BTreeNode    // For leaf nodes (linked list)
}

// BTree represents a B+ tree index
type BTree struct {
	Root  *BTreeNode
	Order int
	mu    sync.RWMutex
}

// NewBTree creates a new B+ tree
func NewBTree() *BTree {
	return &BTree{
		Root: &BTreeNode{
			IsLeaf: true,
			Keys:   make([]int, 0),
			Values: make([]interface{}, 0),
		},
		Order: BTreeOrder,
	}
}

// Insert inserts a key-value pair into the B+ tree
func (t *BTree) Insert(key int, value interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// TODO: Implement B+ tree insertion
	// This is a stub implementation
	return nil
}

// Search searches for a key in the B+ tree
func (t *BTree) Search(key int) (interface{}, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// TODO: Implement B+ tree search
	// This is a stub implementation
	return nil, fmt.Errorf("key not found")
}

// Delete removes a key from the B+ tree
func (t *BTree) Delete(key int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// TODO: Implement B+ tree deletion
	// This is a stub implementation
	return nil
}

// RangeScan performs a range scan on the B+ tree
func (t *BTree) RangeScan(startKey, endKey int) ([]interface{}, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// TODO: Implement range scan
	// This is a stub implementation
	return []interface{}{}, nil
}

// BTreeManager manages multiple B+ tree indexes
type BTreeManager struct {
	indexes map[string]*BTree
	mu      sync.RWMutex
}

// NewBTreeManager creates a new B+ tree manager
func NewBTreeManager() *BTreeManager {
	return &BTreeManager{
		indexes: make(map[string]*BTree),
	}
}

// CreateIndex creates a new index
func (m *BTreeManager) CreateIndex(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.indexes[name]; exists {
		return fmt.Errorf("index %s already exists", name)
	}

	m.indexes[name] = NewBTree()
	return nil
}

// GetIndex retrieves an index by name
func (m *BTreeManager) GetIndex(name string) (*BTree, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	index, exists := m.indexes[name]
	if !exists {
		return nil, fmt.Errorf("index %s does not exist", name)
	}

	return index, nil
}

// DropIndex removes an index
func (m *BTreeManager) DropIndex(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.indexes[name]; !exists {
		return fmt.Errorf("index %s does not exist", name)
	}

	delete(m.indexes, name)
	return nil
}

// ListIndexes returns all index names
func (m *BTreeManager) ListIndexes() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.indexes))
	for name := range m.indexes {
		names = append(names, name)
	}
	return names
}
