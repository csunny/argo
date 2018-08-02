package main

import (

	"flag"
	"github.com/golang/glog"
	mgrpc "github.com/csunny/argo/examples/grpc/grpc_example"
)
func main()  {
	flag.Parse()

	defer glog.Flush()

	if err := mgrpc.Proxy(); err != nil{
		glog.Fatal(err)
	}
}


