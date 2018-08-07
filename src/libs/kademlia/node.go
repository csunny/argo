package kademlia

import "fmt"

func RunServer()  {
	nodeOneId := NewNodeID("FFFFFFFF")
	currentNode := Contract{nodeOneId, "localhost:9999"}
	k := NewKademlia(&currentNode, "NodeOne")
	fmt.Println("server is running at localhost:9999")
	k.Serve()
}


func Client()  {
	nodeOneId := NewNodeID("FFFFFFFF")
	currentNode := Contract{nodeOneId, "localhost:9999"}

	nodeTwoId := NewNodeID("FFFFFFF0")
	other := Contract{nodeTwoId, "localhost:8888"}

	k2 := NewKademlia(&other, "NodeTwo")

	err := k2.Call(
		&currentNode, "KademliaCore.Ping", &PingRequest{RPCHeader{&other, "NodeTwo"}},
		&PingResponce{},
	)

	if err != nil{
		fmt.Println(err)
	}
}