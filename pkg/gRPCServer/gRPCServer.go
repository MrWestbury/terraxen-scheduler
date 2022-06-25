package gRPCServer

import (
	"context"
	"log"

	pb "github.com/MrWestbury/terraxen-scheduler/service"
	"github.com/google/uuid"
)

type RemoteServer struct {
	pb.UnimplementedTerraxenSchedulerServer
	agents *Agentpool
}

func NewRemoteServer() *RemoteServer {
	defaultAgentPool, err := NewAgentpool("default")
	if err != nil {
		log.Printf("failed to create agent pool")
		return nil
	}
	server := &RemoteServer{
		agents: defaultAgentPool,
	}
	return server
}

func (gServ *RemoteServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {

	agentid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("failed to register agent %s: %v", req.Agentname, err)
		return nil, err
	}

	agent := &Agent{
		Id:   agentid.String(),
		Name: req.Agentname,
	}

	gServ.agents.RegisterAgent(agent)

	reply := &pb.RegisterReply{
		Message: "Ok, you are active",
		AgentId: agent.Id,
	}

	return reply, nil
}

func (gServ *RemoteServer) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterReply, error) {
	reply := &pb.UnregisterReply{}
	return reply, nil
}

func (gServer *RemoteServer) Checkin(ctx context.Context, req *pb.CheckinRequest) (*pb.CheckinReply, error) {
	reply := &pb.CheckinReply{}
	return reply, nil
}
