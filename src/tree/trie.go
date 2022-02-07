package tree

import (
	"errors"
)

// A simple trie tree impl

type Action int

const (
	Insert Action = iota
	Update
	Delete
)

type ty int

const (
	unknown ty = iota
	ext
	leaf
	branch
)

type Entry struct {
	action  Action
	key     []byte
	old     []byte
	updated []byte
}

// Node in trie, tree kinds,
// Branch Node [hash_0, hash_1, ...hash_f]
// Extension Node [flag, encodedPath, next hash]
// Leaf Node [flag, encodedPath, value]
type node struct {
	Hash  []byte
	Bytes []byte
	Val   [][]byte
}
type Trie struct {
	rootHash      []byte
	storage       *MemoryStorage
	changelog     []*Entry
	needChangelog bool
}

func (n *node) Type() (ty, error) {
	if n.Val == nil {
		return unknown, errors.New("nil node")
	}

	switch len(n.Val) {
	case 16: // Branch Node
		return branch, nil
	case 3: // Extension Node or Leaf Node
		if n.Val[0] == nil {
			return unknown, errors.New("unknown node type")
		}
		return ty(n.Val[0][0]), nil
	default:
		return unknown, errors.New("wrong node value")
	}
}

func NewTrie(rootHash []byte, storage *MemoryStorage, needChangelog bool) (*Trie, error) {

	t := &Trie{
		rootHash:      rootHash,
		storage:       storage,
		needChangelog: needChangelog,
	}

	if t.rootHash == nil || len(t.rootHash) == 0 {
		return t, nil
	} else if _, err := t.storage.Get(rootHash); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Trie) RootHash() []byte {
	return t.rootHash
}

func (t *Trie) Empty() bool {
	return t.rootHash == nil
}
