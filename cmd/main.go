package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/gRPCServer"
	pb "github.com/MrWestbury/terraxen-scheduler/service"
	"google.golang.org/grpc"
)

func main() {
	portStr := os.Getenv("TERRAXEN_SCHEDULER_GRPC_PORT")
	port := 7100
	if portStr == "" {
		portInt, err := strconv.Atoi(portStr)
		if err == nil {
			port = portInt
		}
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	defaultAgentPool, err := agentpool.NewAgentpool("default")
	if err != nil {
		log.Fatalf("failed to create agent pool")
	}
	handler := gRPCServer.NewRemoteServer(defaultAgentPool)

	pb.RegisterTerraxenSchedulerServer(s, handler)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
