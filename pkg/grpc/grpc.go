package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	helloApi "github.com/mihmatache/hello-world/pkg/api/hello"
	"github.com/mihmatache/hello-world/pkg/greetings"
)


func StartServer(port string, options []grpc.ServerOption) error {
	streamInterceptors := []grpc.StreamServerInterceptor{
		grpc_zap.StreamServerInterceptor(zap.NewExample()),
	}
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(zap.NewExample()),
	}
	options = append(options, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))
	log.Printf("Starting gRPC server on port %s", port)
	s := grpc.NewServer(options...)
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

type Greetings struct{}

func (h Greetings) Hello(ctx context.Context, msg *helloApi.HelloMessage) (*helloApi.HelloMessage, error) {
	log.Printf("%s said hello", msg.Name)
	return Hello(), nil
}

func ClientCall(ctx context.Context, grpcClient *Client) ([]byte, error) {
	conn, err := grpcClient.Build()
	if err != nil {
		return []byte{}, err
	}
	client := helloApi.NewGreetingsClient(conn)
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

func Hello() *helloApi.HelloMessage {
	return &helloApi.HelloMessage{
		Message: "Hello from",
		Name:    greetings.Name,
	}
}

type Client struct {
	addr    string
	options []grpc.DialOption
}

func (c *Client) WithOption(option grpc.DialOption) *Client {
	c.options = append(c.options, option)
	return c
}

func (c *Client) Build() (grpc.ClientConnInterface, error) {
	return grpc.Dial(c.addr, c.options...)
}

func NewClient(addr string) *Client {
	return &Client{
		addr:    addr,
		options: []grpc.DialOption{},
	}
}
