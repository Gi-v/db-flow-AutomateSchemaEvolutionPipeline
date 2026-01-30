package storage

// PageSize defines the size of a database page in bytes
const PageSize = 4096

// Page represents a single page in the storage system
type Page struct {
	ID   int
	Data [PageSize]byte
}

// NewPage creates a new empty page
func NewPage(id int) *Page {
	return &Page{
		ID: id,
	}
}
