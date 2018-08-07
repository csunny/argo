package grpc_example

import (
	"log"
	"net"
	"os"
	"golang.org/x/net/context"
	pb "github.com/csunny/argo/examples/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)

const (
	port        = ":9090"
	defaultName = "Magic"
)

// server
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func RunServer() {
	log.Println("server is running at 127.0.0.1:9090... ")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func Client() {
	// set up a connection to the server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contract the server and print response
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name:name})
	if err != nil{
		log.Fatal("could not greet:", err)
	}
	log.Println("Greeting: ", r.Message)
}


var (
	helloEndPoint = flag.String("hello_endpoint", "localhost:9090", "endpoint of your service")
)
func Proxy() error  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, *helloEndPoint, opts)
	if err != nil{
		return err
	}

	log.Println("The Server is running at 127.0.0.1:8888")
	return http.ListenAndServe(":8888", mux)
}

