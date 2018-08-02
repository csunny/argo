package kademlia

import (
	"testing"
	"fmt"
)

func TestKademliaCore_Ping(t *testing.T) {
	currentNode := Contract{NewRandomNodeID(), "localhost:9999"}

	k := NewKademlia(&currentNode, "Test")
	k.Serve()

	other := Contract{NewRandomNodeID(), "localhost:9999"}
	err := k.Call(
		&other, "KademliaCore.Ping", &PingRequest{RPCHeader{&other, k.NetworkId}},
		&PingResponce{},
	)

	fmt.Println("==================")
	if err != nil{
		t.Error(err)
	}

}

func TestKademliaCore_FindNode(t *testing.T) {
	currentNode := Contract{NewRandomNodeID(), "localhost:9999"}
	k := NewKademlia(&currentNode, "Test")
	kc := KademliaCore{k}

	var contacts [100]Contract

	for i:=0; i<len(contacts); i++{
		contacts[i] = Contract{NewRandomNodeID(), "localhost:9999"}
		err := kc.Ping(&PingRequest{RPCHeader{&contacts[i], k.NetworkId}},
			&PingResponce{},
		)

		if err != nil{
			t.Error(err)
		}
	}

	args := FindNodeRequest{RPCHeader{&contacts[0], k.NetworkId}, contacts[0].id}

	response := FindNodeRequest{}
	err := kc.FindNode(&args, &response)
	if err != nil{
		t.Error(err)
	}

	if len(response.contacts) != BucketSize{
		t.Error()
	}

}
