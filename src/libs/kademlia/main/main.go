package main

import (
	"fmt"
	//"github.com/csunny/argo/src/libs/kademlia"
	"encoding/hex"
)

func main()  {
	//s := func(x int) int {
	//	return x * 2
	//}
	//
	//fmt.Println(s(4))

	//nodeId := kademlia.NewNodeID("node1")
	//fmt.Sprintf("%x", nodeId)
	res, _ := hex.DecodeString("magic")
	fmt.Println(res)
}
