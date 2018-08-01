package kademlia

const BucketSize = 20

type RoutingTable struct {
	node    Contract
	buckets [IdLength * 8]Contract
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
		ret.buckets[i] = *new(Contract)
	}
	ret.node = *node
	return
}

func (table *RoutingTable) Update(contract *Contract) {
	prefix_length := contract.id.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[:prefix_length]

	var element interface{}

	for _, b := range bucket {
		if Equal(b, table) == true {
			element = b
			break
		}
	}

	if element == nil {
		if len(bucket) < BucketSize {
			bucket = append(bucket, *contract)
		}
	}

}

func Equal(x interface{}, table *RoutingTable) bool {
	return x.(*Contract).id.Equals(table.node.id)
}

func (table *RoutingTable) FindClosest(target NodeID, count int) []Contract {

	ret := []Contract{}

	bucket_num := target.Xor(table.node.id).PrefixLen()

	bucket := table.buckets[bucket_num]
	ret = append(ret, bucket)
	for i := 1; (bucket_num-i >= 0 || bucket_num+i < IdLength*8) && len(ret) < count; i++ {
		if bucket_num-i >= 0 {
			bucket = table.buckets[bucket_num-i]
			ret = append(ret, bucket)
		}
		if bucket_num+i < IdLength*8 {
			bucket = table.buckets[bucket_num+i]
			ret = append(ret, bucket)
		}
	}

	if len(ret) > count{
		ret = ret[:count]
	}
	return ret

}
