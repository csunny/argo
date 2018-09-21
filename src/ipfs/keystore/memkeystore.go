package keystore

import ci "github.com/libp2p/go-libp2p-crypto"

// MemKeystore 是内存存储
type MemKeystore struct{
	keys map[string]ci.PrivKey
}

// NewMemKeystore 创建hash表来存储key值
func NewMemKeystore() *MemKeystore{
	return &MemKeystore{make(map[string]ci.PrivKey)}
}

// Has return exists
// TODO need to implement interface api 