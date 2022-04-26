package repositories

// Type defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

// Storage defines the functionality of a data store for the beer service.
type Storage interface {
}

// DB is the "global" storage instance
var DB Storage

func NewStorage(t Type) *MemoryStorage {
	switch t {
	case Memory:
		return new(MemoryStorage)

	case JSON:
		// for the moment storage location for JSON files is the current working directory
	}

	return nil
}
