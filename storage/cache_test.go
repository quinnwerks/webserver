package storage

import (
	"testing"
)

func TestCachePutSuccess(t *testing.T) {

}

func TestCachePutError(t *testing.T) {

}

func TestCacheGetSuccess(t *testing.T) {

}

func TestCacheGetError(t *testing.T) {

}

func TestCacheNodeSize(t *testing.T) {
	key := "four"
	value := "thirteenbytes"
	golden_size := 17
	node := newCacheNode(key, value)
	size := node.Size()
	if size != golden_size {
		t.Errorf("(golden) %d != (actual) %d", golden_size, size)
	}
}

func TestCacheNodeHash(t *testing.T) {
}
