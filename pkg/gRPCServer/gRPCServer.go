package gRPCServer

import (
	"context"
	"log"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	pb "github.com/MrWestbury/terraxen-scheduler/service"
	"github.com/google/uuid"
)

type RemoteServer struct {
	pb.UnimplementedTerraxenSchedulerServer
	agents *agentpool.Agentpool
}

func NewRemoteServer(agentpool *agentpool.Agentpool) *RemoteServer {

	server := &RemoteServer{
		agents: agentpool,
	}
	return server
}

func (gServ *RemoteServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {

	agentid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("failed to register agent %s: %v", req.Agentname, err)
		return nil, err
	}

	agent := &agentpool.Agent{
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
	gServ.agents.UnregisterAgent(req.GetAgentId())
	reply := &pb.UnregisterReply{
		Message: "Agent unregistered",
	}
	return reply, nil
}

func (gServer *RemoteServer) Checkin(ctx context.Context, req *pb.CheckinRequest) (*pb.CheckinReply, error) {
	reply := &pb.CheckinReply{}
	return reply, nil
}

func (gServer *RemoteServer) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobReply, error) {
	reply := &pb.GetJobReply{}
	return reply, nil
}

func (gServer *RemoteServer) UpdateJob(ctx context.Context, req *pb.UpdateJobStateRequest) (*pb.UpdateJobStateReply, error) {

}
