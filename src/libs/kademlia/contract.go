package kademlia

import (
	"fmt"
)

type Contract struct {
	Id      NodeID
	Address string
}

func (contract *Contract) String() string {
	return fmt.Sprintf("Contract(\"%s\", \"%s\")", contract.Id, contract.Address)
}


