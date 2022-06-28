package grpc_server

import (
	"context"
	"errors"
	"log"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/jobpool"
	pb "github.com/MrWestbury/terraxen-scheduler/service"
	"github.com/google/uuid"
)

type RemoteServerHandler struct {
	pb.UnimplementedTerraxenSchedulerServer
	agents *agentpool.Agentpool
	jobs   jobpool.JobPoolInterface
}

func NewRemoteServerHandler(agentpool *agentpool.Agentpool, jobpool jobpool.JobPoolInterface) *RemoteServerHandler {

	server := &RemoteServerHandler{
		agents: agentpool,
		jobs:   jobpool,
	}
	return server
}

func (gServ *RemoteServerHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	log.Printf("registering new agent: %s", req.Agentname)
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

func (gServ *RemoteServerHandler) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterReply, error) {
	log.Printf("Unregistering agent %s", req.GetAgentId())
	gServ.agents.UnregisterAgent(req.GetAgentId())
	reply := &pb.UnregisterReply{
		Message: "Agent unregistered",
	}
	return reply, nil
}

func (gServer *RemoteServerHandler) Checkin(ctx context.Context, req *pb.CheckinRequest) (*pb.CheckinReply, error) {
	log.Printf("Agent checking in %s", req.GetAgentId())

	reply := &pb.CheckinReply{
		JobId: "",
	}

	isJob := gServer.jobs.GetANewJob()
	if isJob != nil {
		success := gServer.jobs.AllocateToAgent(isJob.Id, req.GetAgentId())
		if success {
			reply.JobId = isJob.Id
		}
	}
	return reply, nil
}

func (gServer *RemoteServerHandler) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobReply, error) {
	log.Printf("Agent %s requesting job details %s", req.GetAgentId(), req.GetJobId())
	job, err := gServer.jobs.GetJobById(req.GetJobId())
	if err != nil {
		log.Printf("failed to get job: %v", err)
		return nil, err
	}

	if job.AgentAllocated != req.GetAgentId() {
		log.Printf("wrong agent getting job")
		return nil, errors.New("job not allocated to agent")
	}
	reply := &pb.GetJobReply{
		JobId:  job.Id,
		GitUrl: job.GitUrl,
	}
	return reply, nil
}

func (gServer *RemoteServerHandler) UpdateJob(ctx context.Context, req *pb.UpdateJobStateRequest) (*pb.UpdateJobStateReply, error) {
	log.Printf("Updating job details %s", req.GetJobId())
	reply := &pb.UpdateJobStateReply{
		Message: "Ok",
	}
	return reply, nil
}
