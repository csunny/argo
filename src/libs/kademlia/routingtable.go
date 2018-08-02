package kademlia

import (
	"container/list"
	"fmt"
)

const BucketSize = 20

type RoutingTable struct {
	node    Contract
	buckets [IdLength * 8]*list.List
}

type ContractRecord struct {
	node    *Contract
	sortKey NodeID
}

func (rec *ContractRecord) Less(other interface{}) bool {
	return rec.sortKey.Less(other.(*ContractRecord).sortKey)
}

func NewRoutingTable(node *Contract) (ret *RoutingTable) {
	ret = new(RoutingTable)
	for i := 0; i < IdLength*8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.node = *node
	return
}

func (table *RoutingTable) Update(contract *Contract) {
	prefixLength := contract.id.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[prefixLength]

	var element interface{}

	iter := bucket.Front()

	for iter != nil {
		if Equal(iter.Value, table) == true {
			element = bucket.Front()
		} else {
			iter = bucket.Front().Prev()
		}
	}

	iter_back := bucket.Back()
	for iter_back != nil {
		if Equal(iter_back.Value, table) == true {
			element = bucket.Back()
		} else {
			iter_back = bucket.Back().Next()
		}
	}

	if element == nil {
		fmt.Println("-------")
		if bucket.Len() < BucketSize {
			bucket.PushFront(contract)
		}
		fmt.Println(bucket.Front().Value)
		// 剔除节点
		// todo list

	} else {
		bucket.MoveToFront(element.(*list.Element))
	}

}

func Equal(x interface{}, table *RoutingTable) bool {
	return x.(*Contract).id.Equals(table.node.id)
}

func (table *RoutingTable) FindClosest(target NodeID, count int) (ret []interface{}) {

	bucket_num := target.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[bucket_num]

	ret = append(ret, bucket.Front())

	for i := 1; (bucket_num-i >= 0 || bucket_num+i < IdLength*8) && len(ret) < count; i++ {
		if bucket_num-i >= 0 {
			bucket = table.buckets[bucket_num-i]
			ret = append(ret, bucket.Front())
		}
		if bucket_num+i < IdLength*8 {
			bucket = table.buckets[bucket_num+i]
			ret = append(ret, bucket.Front())
		}
	}

	if len(ret) > count {
		ret = ret[:count]
	}
	return
}
