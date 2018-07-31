package kademlia

import (
	"container/list"
	"github.com/csunny/argo/src/iterator"

)

const BucketSize = 20

type RoutingTable struct {
	node    Contract
	buckets [IdLength * 8]*list.List
}

type ContractRecord struct {
	node *Contract
	sortKey NodeID
}

func NewRoutingTable(node *Contract) (ret *RoutingTable)  {
	ret = new(RoutingTable)
	for i:=0; i<IdLength * 8; i++{
		ret.buckets[i] = list.New()
	}
	ret.node = *node
	return
}

func (table *RoutingTable) Update(contract *Contract)  {
	prefix_length := contract.id.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[prefix_length]

	//element := iterable.Find(bucket, func(x interface{}) bool {
	//	return x.(*Contract).id.Equals(table.node.id)
	//})

}

//func copyToVector(start, end *list.Element, vec *)  {
//
//}
