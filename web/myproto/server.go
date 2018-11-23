package myproto

import (
	"fmt"
	"log"
	"net"

	"github.com/IhorBondartsov/datasaver/database"
	"github.com/IhorBondartsov/datasaver/entity"
	"github.com/IhorBondartsov/datasaver/web/myproto/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCServerCfg - config for proto server
type GRPCServerCfg struct {
	Port string
	DB   database.DataBase
}

// NewServer create new server
func NewServer(cfg GRPCServerCfg) *server {
	return &server{
		Port: cfg.Port,
		db:   cfg.DB,
	}
}

// server represents the gRPC server
type server struct {
	Port string
	db   database.DataBase
}

// Start - start work server via TCP
func (s *server) Start() error {
	// create a listener on TCP port

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

// Print - print all information from database into log
func (s *server) Print(ctx context.Context, in *pb.Nothing) (*pb.Nothing, error) {
	fmt.Println("Print")
	s.db.Print()
	return nil, nil
}

// Save save PersonData to database
func (s *server) Save(ctx context.Context, in *pb.PersonData) (*pb.Nothing, error) {
	fmt.Println("Save")
	err := s.db.Save(convertPBPersonDataToAPI(in))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.Nothing{}, nil
}

func convertPBPersonDataToAPI(in *pb.PersonData) entity.PersonData {
	return entity.PersonData{
		Email:        in.Email,
		Id:           int(in.ID),
		MobileNumber: in.MobileNumber,
		Name:         in.Name,
	}
}
