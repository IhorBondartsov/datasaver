package myproto

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/IhorBondartsov/datasaver/web/myproto/pb"
)

// GRPCServerCfg - config for proto server
type GRPCServerCfg struct {
	Port string
}

// NewServer create new server
func NewServer(cfg GRPCServerCfg) *server {
	return &server{
		Port: cfg.Port,
	}
}

// server represents the gRPC server
type server struct {
	Port string
}

// Start - start work server via TCP
func (s *server) Start() error {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
  // attach the Ping service to the server
  pb.RegisterCSVSenderServer(grpcServer, s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return nil
}

// SayHello generates response to a Ping request
func (s *server) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &pb.PingMessage{Greeting: fmt.Sprintf("Hello. I see your message: %s ", in.Greeting)}, nil
}

// SayHello generates response to a Ping request
func (s *server) Print(ctx context.Context, in *pb.Nothing) (*pb.Nothing, error) {
	
	return &pb.Nothing{}, nil
}

// SayHello generates response to a Ping request
func (s *server) Save(ctx context.Context, in *pb.PersonData) (*pb.Nothing, error) {
	
	return &pb.Nothing{}, nil
}