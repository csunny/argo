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
	prefixLength := contract.Id.Xor(table.node.Id).PrefixLen()
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

	iterBack := bucket.Back()
	for iterBack != nil {
		if Equal(iterBack.Value, table) == true {
			element = bucket.Back()
		} else {
			iterBack = bucket.Back().Next()
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
	return x.(*Contract).Id.Equals(table.node.Id)
}

func (table *RoutingTable) FindClosest(target NodeID, count int) (ret []interface{}) {

	bucketNum := target.Xor(table.node.Id).PrefixLen()
	bucket := table.buckets[bucketNum]

	ret = append(ret, bucket.Front())

	for i := 1; (bucketNum-i >= 0 || bucketNum+i < IdLength*8) && len(ret) < count; i++ {
		if bucketNum-i >= 0 {
			bucket = table.buckets[bucketNum-i]
			ret = append(ret, bucket.Front())
		}
		if bucketNum+i < IdLength*8 {
			bucket = table.buckets[bucketNum+i]
			ret = append(ret, bucket.Front())
		}
	}

	if len(ret) > count {
		ret = ret[:count]
	}
	return
}
