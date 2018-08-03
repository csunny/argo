package kademlia

import (
	"net/rpc"
	"net"
	"fmt"
	//"net/http"
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

// rpc server
func (k *Kademlia) Serve() error {
	rpc.RegisterName("KademliaCore", &KademliaCore{k})

	listener, err := net.Listen("tcp", k.routers.node.Address)

	//rpc.HandleHTTP()

	if err != nil {
		return err
	}

	conn, err := listener.Accept()
	fmt.Println(conn.LocalAddr())
	if err != nil {
		return err
	}
	//go http.Serve(listener, nil)
	go rpc.ServeConn(conn)

	go fmt.Println(k.routers)
	select{}
	return nil
}

// rpc client
func (k *Kademlia) Call(contract *Contract, method string, args, reply interface{}) error {
	client, err := rpc.Dial("tcp", contract.Address)
	if err != nil {
		return err
	}

	fmt.Println(method, args)
	err = client.Call(method, args, reply)
	if err != nil {
		return err
	}

	fmt.Println("reply:", reply)

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
	//if request.networkId != k.NetworkId {
	//	return fmt.Errorf("Excepted networkID %s, got %s", k.NetworkId, request.networkId)
	//}

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

	fmt.Printf("ping from %s", args.RPCHeader)
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
