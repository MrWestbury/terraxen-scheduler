package grpc_server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/jobpool"
	pb "github.com/MrWestbury/terraxen-scheduler/service"
	"google.golang.org/grpc"
)

type RemoteServer struct {
	listen_host      string
	port             int
	defaultAgentPool *agentpool.Agentpool
	jobPool          jobpool.JobPoolInterface
}

func NewRemoteServer(listener_host string, port int, defaultAgentPool *agentpool.Agentpool, jobPool jobpool.JobPoolInterface) *RemoteServer {
	rs := &RemoteServer{
		listen_host:      listener_host,
		port:             port,
		defaultAgentPool: defaultAgentPool,
		jobPool:          jobPool,
	}

	return rs
}

func (rs *RemoteServer) Start(wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", rs.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	handler := NewRemoteServerHandler(rs.defaultAgentPool, rs.jobPool)

	pb.RegisterTerraxenSchedulerServer(s, handler)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	wg.Done()
}
