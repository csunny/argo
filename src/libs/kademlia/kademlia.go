package kademlia

import (
	"container/heap"
)

type Kademlia struct {
	routers *RoutingTable
	NetworkId string
}

func NewKademlia(contract *Contract, networkId string) (ret *Kademlia)  {
	ret = new(Kademlia)
	return
}

