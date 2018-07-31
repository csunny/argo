package kademlia

import (
	"fmt"
)

type Contract struct {
	id      NodeID
	address string
}

func (contract *Contract) String() string {
	return fmt.Sprintf("Contract(\"s\", \"s\")", contract.id, contract.address)
}


