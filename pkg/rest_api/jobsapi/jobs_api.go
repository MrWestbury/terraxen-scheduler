package jobsapi

import (
	"net/http"

	"github.com/MrWestbury/terraxen-scheduler/pkg/jobpool"
	"github.com/gin-gonic/gin"
)

type JobsApi struct {
	jobpool jobpool.JobPoolInterface
}

func NewJobsApi(parentGroup *gin.RouterGroup, jobPool jobpool.JobPoolInterface) *JobsApi {
	ja := &JobsApi{
		jobpool: jobPool,
	}

	jobsRouter := parentGroup.Group("jobs")

	jobsRouter.GET("/", ja.ListJobs)   // List jobs
	jobsRouter.POST("/", ja.CreateJob) // Create a job

	jobsRouter.GET("/:jobid") // Get details of a job
	jobsRouter.PUT("/:jobid") // Update a job

	return ja
}

// ListJobs handles listing of jobs
func (ja *JobsApi) ListJobs(c *gin.Context) {
	jobs, _ := ja.jobpool.ListJobs()

	c.IndentedJSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"limit":  25,
			"offset": 0,
			"count":  len(jobs),
		},
		"data": jobs,
	})
}

func (ja *JobsApi) CreateJob(c *gin.Context) {

	var requestBody NewJobRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newJob := jobpool.NewJob(requestBody.GitUrl, requestBody.TerraformVersion)
	ja.jobpool.AddJob(newJob)
	c.Status(http.StatusCreated)
}
