package jobpool

import "github.com/google/uuid"

const (
	JOB_STATUS_NEW       = "new"
	JOB_STATUS_LOCKED    = "locked"
	JOB_STATUS_ALLOCATED = "allocated"
	JOB_STATUS_RUNNING   = "running"
	JOB_STATUS_ERRORED   = "errored"
	JOB_STATUS_SUCCESS   = "success"
	JOB_STATUS_UNKNOWN   = "unknown"
)

type Job struct {
	Id               string
	GitUrl           string
	TerraformVersion string
	Status           string
	AgentAllocated   string
	Logs             []string
}

func NewJob(url string, tfVersion string) *Job {
	jobId := uuid.NewString()
	j := &Job{
		Id:               jobId,
		GitUrl:           url,
		TerraformVersion: tfVersion,
		Status:           JOB_STATUS_NEW,
		AgentAllocated:   "",
		Logs:             make([]string, 0),
	}

	return j
}

func (j *Job) Log(message string) {
	j.Logs = append(j.Logs, message)
}
