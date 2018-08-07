package rpc_example

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
)

type Greeter struct {
	
}

func (p *Greeter) Greet(request string, response *string) error  {
	*response = "Hello: " + request
	return nil
}

func Server()  {
	rpc.RegisterName("Greeter", new(Greeter))
	listener, err := net.Listen("tcp", ":8888")

	fmt.Println("Server is running at localhost:8888...")
	if err != nil{
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil{
		log.Fatal("Accept error", err)
	}

	rpc.ServeConn(conn)
}

func Client()  {
	client, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil{
		log.Fatal("dialing:", err)
	}

	var response string
	err = client.Call("Greeter.Greet", "magic", &response)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(response)
}