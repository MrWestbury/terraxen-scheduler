package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/grpc_server"
	"github.com/MrWestbury/terraxen-scheduler/pkg/jobpool"
	restapi "github.com/MrWestbury/terraxen-scheduler/pkg/rest_api"
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
	restPortStr := os.Getenv("TERRAXEN_SCHEDULER_REST_PORT")
	restPort := 7101
	if restPortStr == "" {
		restPortInt, err := strconv.Atoi(restPortStr)
		if err == nil {
			restPort = restPortInt
		}
	}

	jobs := jobpool.NewJobPool()

	defaultAgentPool, err := agentpool.NewAgentpool("default")
	if err != nil {
		log.Fatalf("failed to create agent pool")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	rpc_server := grpc_server.NewRemoteServer("", port, defaultAgentPool, jobs)
	go rpc_server.Start(&wg)

	// Rest server
	rest_server := restapi.NewRestApiServer("", restPort, defaultAgentPool, jobs)
	go rest_server.Start(&wg)

	wg.Wait()
}
