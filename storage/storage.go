package server

// Storage is an interface which can be used to store and
// retrieve data from disk or memory using a NoSQL convention.
// Storage interfaces should be thread safe at the interface level.
type Storage interface {
	// Put a value into storage
	Put(key string, value string) error
	// Get a value from storage
	Get(key string) (StorageNode, error)
}

// StorageNode is the Storage types representation of data.
type StorageNode interface {
	GetHash() int
	GetSize() int
}
