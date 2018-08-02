package kademlia

import (
	"testing"
	"fmt"
)

func TestNodeID(t *testing.T)  {
	a := NewNodeID("Node1")
	b := NewNodeID("Node2")
	c := NewNodeID("Node3")

	fmt.Printf("%s %s %s\n", a, b, c)

	d := NewRandomNodeID()
	fmt.Println(d)

	if !a.Equals(a){
		t.Fatal("Equals func is wrong!")
	}

	if a.Equals(b){
		t.Fatal("Equals func is wrong!")
	}

	e := NodeID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	f := NodeID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 18}
	g := NodeID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}

	if !e.Xor(f).Equals(g){
		t.Error(a.Xor(b))
	}

	if g.PrefixLen() != 151{
		t.Error("prefixlen is error")
	}

	if b.Less(a){
		t.Error("b should more a!")
	}

	fmt.Println(e, e.PrefixLen())

}