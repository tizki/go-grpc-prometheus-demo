package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"go-grpc-prometheus-demo/config"
	"go-grpc-prometheus-demo/routeguide"
	pb "go-grpc-prometheus-demo/routeguide"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func (s *routeGuideServer) SayHello(ctx context.Context, message *pb.HelloRequest) (*pb.HelloReply, error){
	log.Printf("Received: %v", message.GetName())
	return &pb.HelloReply{Message: "Hello " + message.GetName()}, nil
}

func (s *routeGuideServer) EchoService(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	return &pb.Message{Message: fmt.Sprintf("echo %s", msg.Message)}, nil
}

func main() {

	// Create the main listener.
	s, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GetServerPort()))
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(s)

	// Match connections in order:
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// gRPC server
	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	routeguide.RegisterRouteGuideServer(grpcS, newServer())
	// prometheus metrics server
	grpc_prometheus.Register(grpcS)
	httpS := &http.Server{
		Handler: promhttp.Handler(),
	}

	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	m.Serve()
}

type routeGuideServer struct {
	routeguide.UnimplementedRouteGuideServer
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}