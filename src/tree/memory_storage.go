package tree

import (
	"encoding/hex"
	"fmt"
	"sync"
)

// MemoryStorage the nodes in trie
type MemoryStorage struct {
	data *sync.Map
}

func NewMemoryStorage() (*MemoryStorage, error) {
	return &MemoryStorage{
		data: new(sync.Map),
	}, nil
}

// Get return value to the key in Storage
func (db *MemoryStorage) Get(key []byte) ([]byte, error) {
	if entry, ok := db.data.Load(hex.EncodeToString(key)); ok {
		return entry.([]byte), nil
	}
	return nil, fmt.Errorf("key not found: %s", hex.EncodeToString(key))
}

// Put
func (db *MemoryStorage) Put(key []byte, value []byte) error {
	db.data.Store(hex.EncodeToString(key), value)
	return nil
}

func (db *MemoryStorage) Delete(key []byte) error {
	db.data.Delete(hex.EncodeToString(key))
	return nil
}
