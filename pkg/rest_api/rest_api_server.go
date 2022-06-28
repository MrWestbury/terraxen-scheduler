package restapi

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/MrWestbury/terraxen-scheduler/pkg/agentpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/jobpool"
	"github.com/MrWestbury/terraxen-scheduler/pkg/rest_api/agentsapi"
	"github.com/MrWestbury/terraxen-scheduler/pkg/rest_api/jobsapi"
	"github.com/gin-gonic/gin"
)

type RestApiServer struct {
	listenerHost string
	port         int
	agentPoolSvc *agentpool.Agentpool
	jobPool      jobpool.JobPoolInterface
}

func NewRestApiServer(listenerHostname string, port int, agentPool *agentpool.Agentpool, jobPool jobpool.JobPoolInterface) *RestApiServer {
	ras := &RestApiServer{
		listenerHost: listenerHostname,
		port:         port,
		agentPoolSvc: agentPool,
		jobPool:      jobPool,
	}

	return ras
}

func (ras *RestApiServer) Start(wg *sync.WaitGroup) {
	router := gin.Default()

	apiGroup := router.Group("api")
	v1Group := apiGroup.Group("v1")

	v1Group.GET("/", func(c *gin.Context) {
		resp := gin.H{
			"message": "ok",
		}

		c.IndentedJSON(http.StatusOK, resp)
	})

	agentsapi.NewAgentPoolApi(v1Group)
	jobsapi.NewJobsApi(v1Group, ras.jobPool)

	listenAddr := fmt.Sprintf("%s:%d", ras.listenerHost, ras.port)
	log.Printf("Rest API listening on %s", listenAddr)
	err := router.Run(listenAddr)
	if err != nil {
		log.Printf("failed to start rest server: %v", err)
	}
	wg.Done()
}
