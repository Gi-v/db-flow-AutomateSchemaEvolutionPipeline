package transaction

import (
	"fmt"
	"sync"
	"time"
)

// LockType represents the type of lock
type LockType int

const (
	ReadLock LockType = iota
	WriteLock
)

// Lock represents a lock on a resource
type Lock struct {
	Resource  string
	Type      LockType
	HolderID  uint64
	Timestamp time.Time
}

// LockManager manages locks for concurrency control
type LockManager struct {
	locks      map[string][]Lock
	waitQueue  map[string][]Lock
	mu         sync.Mutex
	nextLockID uint64
}

// NewLockManager creates a new lock manager
func NewLockManager() *LockManager {
	return &LockManager{
		locks:      make(map[string][]Lock),
		waitQueue:  make(map[string][]Lock),
		nextLockID: 1,
	}
}

// AcquireReadLock acquires a read lock on a resource
func (lm *LockManager) AcquireReadLock(resource string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	// Check if there are any write locks on the resource
	locks := lm.locks[resource]
	for _, lock := range locks {
		if lock.Type == WriteLock {
			// Wait for write lock to be released (simplified - would normally use condition variables)
			return fmt.Errorf("write lock exists on resource %s", resource)
		}
	}

	// Grant read lock
	lockID := lm.nextLockID
	lm.nextLockID++

	lock := Lock{
		Resource:  resource,
		Type:      ReadLock,
		HolderID:  lockID,
		Timestamp: time.Now(),
	}

	lm.locks[resource] = append(lm.locks[resource], lock)
	return nil
}

// AcquireWriteLock acquires a write lock on a resource
func (lm *LockManager) AcquireWriteLock(resource string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	// Check if there are any locks on the resource
	locks := lm.locks[resource]
	if len(locks) > 0 {
		// Wait for all locks to be released (simplified)
		return fmt.Errorf("locks exist on resource %s", resource)
	}

	// Grant write lock
	lockID := lm.nextLockID
	lm.nextLockID++

	lock := Lock{
		Resource:  resource,
		Type:      WriteLock,
		HolderID:  lockID,
		Timestamp: time.Now(),
	}

	lm.locks[resource] = append(lm.locks[resource], lock)
	return nil
}

// ReleaseReadLock releases a read lock
func (lm *LockManager) ReleaseReadLock(resource string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	locks := lm.locks[resource]
	if len(locks) == 0 {
		return fmt.Errorf("no locks on resource %s", resource)
	}

	// Remove first read lock (simplified)
	for i, lock := range locks {
		if lock.Type == ReadLock {
			lm.locks[resource] = append(locks[:i], locks[i+1:]...)
			break
		}
	}

	// Clean up empty entries
	if len(lm.locks[resource]) == 0 {
		delete(lm.locks, resource)
	}

	return nil
}

// ReleaseWriteLock releases a write lock
func (lm *LockManager) ReleaseWriteLock(resource string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	locks := lm.locks[resource]
	if len(locks) == 0 {
		return fmt.Errorf("no locks on resource %s", resource)
	}

	// Remove first write lock (simplified)
	for i, lock := range locks {
		if lock.Type == WriteLock {
			lm.locks[resource] = append(locks[:i], locks[i+1:]...)
			break
		}
	}

	// Clean up empty entries
	if len(lm.locks[resource]) == 0 {
		delete(lm.locks, resource)
	}

	return nil
}

// DetectDeadlock detects deadlocks in the system
func (lm *LockManager) DetectDeadlock() (bool, error) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	// TODO: Implement deadlock detection using wait-for graph
	// This is a stub implementation
	return false, nil
}

// GetLocks returns all current locks (for testing/debugging)
func (lm *LockManager) GetLocks() map[string][]Lock {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	locks := make(map[string][]Lock)
	for k, v := range lm.locks {
		locks[k] = append([]Lock{}, v...)
	}
	return locks
}
