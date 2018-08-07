package kademlia

import (
	"testing"
	"fmt"
)

func TestNewRoutingTable(t *testing.T) {
	n1 := NewNodeID("FFFFFFFF")
	n2 := NewNodeID("FFFFFFF0")
	n3 := NewNodeID("11111111")

	rt := NewRoutingTable(&Contract{n1, "localhost:5000"})
	rt.Update(&Contract{n2, "localhost:5001"})
	rt.Update(&Contract{n3, "localhost:5002"})

	//fmt.Println(rt)
	n4 := NewNodeID("22222222")
	res := rt.FindClosest(n4, 2)

	for _, r := range res{
		fmt.Println(r)
	}

}