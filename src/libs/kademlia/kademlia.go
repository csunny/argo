package kademlia

import (
	"net/rpc"
	"net"
	"net/http"
	"fmt"
	"log"
)

type Kademlia struct {
	routers   *RoutingTable
	NetworkId string
}

type KademliaCore struct {
	kad *Kademlia
}

type RPCHeader struct {
	Sender    *Contract
	networkId string
}
type PingRequest struct {
	RPCHeader
}

type PingResponce struct {
	RPCHeader
}

type FindNodeRequest struct {
	RPCHeader
	target NodeID
}

type FindNodeResponse struct {
	RPCHeader
	contacts []Contract
}

func NewKademlia(contract *Contract, networkId string) (ret *Kademlia) {
	ret = new(Kademlia)
	ret.routers = NewRoutingTable(contract)
	ret.NetworkId = networkId
	return
}

func (k *Kademlia) Serve() error {
	rpc.Register(&KademliaCore{k})

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", k.routers.node.address)

	if err != nil {
		return err
	}
	go http.Serve(l, nil)
	return nil
}

func (k *Kademlia) Call(contract *Contract, method string, args, reply interface{}) error {
	client, err := rpc.DialHTTP("tcp", contract.address)
	if err != nil {
		return err
	}

	err = client.Call(method, args, reply)
	if err != nil {
		return nil
	}

	k.routers.Update(contract)
	return nil
}

func (k *Kademlia) sendQuery(node *Contract, target NodeID, done chan []Contract) {
	args := FindNodeRequest{RPCHeader{&k.routers.node, k.NetworkId}, target}
	reply := FindNodeResponse{}

	err := k.Call(node, "KademliaCore.FindNode", &args, &reply)
	if err != nil {
		done <- []Contract{}
	}

	done <- reply.contacts

}

func (k *Kademlia) InterativeFindNode(target NodeID, delta int) []Contract {
	return []Contract{}
}

func (k *Kademlia) HandleRPC(request, response *RPCHeader) error {
	if request.networkId != k.NetworkId {
		return fmt.Errorf("---")
	}

	if request.Sender != nil {
		k.routers.Update(request.Sender)
	}

	response.Sender = &k.routers.node
	return nil
}

func (kc *KademliaCore) Ping(args *PingRequest, response *PingResponce) error {
	err := kc.kad.HandleRPC(&args.RPCHeader, &response.RPCHeader)
	if err != nil {
		return err
	}

	log.Printf("ping from %s", args.RPCHeader)
	return nil
}

func (kc *KademliaCore) FindNode(args *FindNodeRequest, response *FindNodeResponse) error {
	err := kc.kad.HandleRPC(&args.RPCHeader, &response.RPCHeader)

	if err != nil {
		return err
	}
	contancts := kc.kad.routers.FindClosest(args.target, BucketSize)
	response.contacts = make([]Contract, len(contancts))

	for i := 0; i < len(contancts); i++ {
		response.contacts[i] = *contancts[i].(*ContractRecord).node
	}

	return nil
}
