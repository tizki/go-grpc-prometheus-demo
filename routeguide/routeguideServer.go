package routeguide

import (
	"context"
	"fmt"
	pb "github.com/nsoushi/go-grpc-prometheus-demo/protobuf"
	"log"
)

type routeGuideServer struct {
}

func (s *routeGuideServer) SayHello(ctx context.Context, message *HelloRequest) (*HelloReply, error){
	log.Printf("Received: %v", message.GetName())
	return &HelloReply{Message: "Hello " + message.GetName()}, nil
}

func (s *routeGuideServer) EchoService(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	return &pb.Message{Message: fmt.Sprintf("echo %s", msg.Message)}, nil
}