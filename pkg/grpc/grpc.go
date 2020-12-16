package grpc

import(
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	helloApi "github.com/mihmatache/hello-world/pkg/api/hello"
	"github.com/mihmatache/hello-world/pkg/greetings"


)

func StartServer(port string) error{
	log.Printf("Starting gRPC server on port %s", port)
	s := grpc.NewServer()
	defer s.GracefulStop()
	helloApi.RegisterGreetingsServer(s, Greetings{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

type Greetings struct {}

func (h Greetings) Hello(ctx context.Context, msg *helloApi.HelloMessage) (*helloApi.HelloMessage, error) {
	log.Printf("%s said hello", msg.Name)
	return Hello(), nil
}

func ClientCall(address, port string, timeout int) ([]byte, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := helloApi.NewGreetingsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)
	defer cancel()
	reply, err := client.Hello(ctx, Hello())
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(reply)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Hello() *helloApi.HelloMessage{
	return &helloApi.HelloMessage{
		Message: "Hello from",
		Name: greetings.Name,
	}
}