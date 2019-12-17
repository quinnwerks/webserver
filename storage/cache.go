package storage

import ()

// Cache is a struct which implements the storage interface.
// A cache in this case stores data in memory.
// Will be used transparently speed up disc access.
type Cache struct {
	p policy
}

// Put a value into the cache
func (c *Cache) Put() {

}

// Get a value from the cache
func (c *Cache) Get() {

}

type policy interface {
}

type cacheNode struct {
	key   string
	value string
}

func newCacheNode(key string, value string) *cacheNode {
	return &cacheNode{key: key, value: value}
}

func (node *cacheNode) Size() int {
	return len(node.key) + len(node.value)
}

func (node *cacheNode) Hash() string {
	return ""
}
