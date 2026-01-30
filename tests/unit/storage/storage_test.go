package storage

import (
	"testing"

	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPage(t *testing.T) {
	page := storage.NewPage(1)
	assert.NotNil(t, page)
	assert.Equal(t, storage.PageID(1), page.ID)
	assert.Equal(t, storage.PageSize, len(page.Data))
	assert.False(t, page.Dirty)
	assert.Equal(t, 0, page.PinCount)
}

func TestPagePinUnpin(t *testing.T) {
	page := storage.NewPage(1)

	// Test pin
	page.Pin()
	assert.Equal(t, 1, page.PinCount)
	assert.True(t, page.IsPinned())

	// Test multiple pins
	page.Pin()
	assert.Equal(t, 2, page.PinCount)

	// Test unpin
	err := page.Unpin()
	assert.NoError(t, err)
	assert.Equal(t, 1, page.PinCount)

	// Unpin again
	err = page.Unpin()
	assert.NoError(t, err)
	assert.Equal(t, 0, page.PinCount)
	assert.False(t, page.IsPinned())

	// Test unpin error
	err = page.Unpin()
	assert.Error(t, err)
}

func TestPageWriteRead(t *testing.T) {
	page := storage.NewPage(1)

	data := []byte("Hello, World!")
	err := page.Write(0, data)
	assert.NoError(t, err)
	assert.True(t, page.IsDirty())

	readData, err := page.Read(0, len(data))
	assert.NoError(t, err)
	assert.Equal(t, data, readData)
}

func TestBufferPool(t *testing.T) {
	bp := storage.NewBufferPool(10)
	assert.NotNil(t, bp)
	assert.Equal(t, 0, bp.Size())

	// Fetch a page
	page, err := bp.FetchPage(1)
	assert.NoError(t, err)
	assert.NotNil(t, page)
	assert.Equal(t, 1, bp.Size())

	// Unpin the page
	err = bp.UnpinPage(1, false)
	assert.NoError(t, err)
}

func TestTableManager(t *testing.T) {
	bp := storage.NewBufferPool(100)
	tm := storage.NewTableManager(bp)

	// Create table
	schema := storage.Schema{
		Columns: []storage.Column{
			{Name: "id", Type: "int"},
			{Name: "name", Type: "string"},
		},
	}
	err := tm.CreateTable("users", schema)
	assert.NoError(t, err)

	// Get table
	table, err := tm.GetTable("users")
	assert.NoError(t, err)
	assert.Equal(t, "users", table.Name)

	// List tables
	tables := tm.ListTables()
	assert.Contains(t, tables, "users")

	// Drop table
	err = tm.DropTable("users")
	assert.NoError(t, err)
}
