package kademlia

import (
	"encoding/hex"
	"math/rand"
)

const IdLength = 20

type NodeID [IdLength]byte

func NewNodeID(data string) (ret NodeID) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}

	return
}

func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}

	return
}

func (node NodeID) String() string {
	return hex.EncodeToString(node[:])
}

func (node NodeID) Equals(other NodeID) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other.(NodeID)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

func (node NodeID) Xor(other NodeID) (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = node[i] ^ other[i]
	}
	return
}

func (node NodeID) PrefixLen() (ret int) {
	for i := 0; i < IdLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return IdLength*8 - 1
}


