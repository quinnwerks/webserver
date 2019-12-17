package storage

// Storage is an interface which can be used to store and
// retrieve data from disk or memory using a NoSQL convention.
// Storage interfaces public functions will be thread safe.
type Storage interface {
	// Put a value into storage
	Put(key string, value string) error
	// Get a value from storage
	Get(key string) (StorageNode, error)
}

// StorageNode is the Storage types representation of data.
type StorageNode interface {
	// Gets the hash of the storage nodes key.
	GetHash() string
	// Returns the size of the node in bytes.
	GetSize() int
}
