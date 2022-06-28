// Package jobpool provides a structure for managing a pool of jobs
package jobpool

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrJobDoesNotExist = errors.New("job does not exist")
)

type JobPoolInterface interface {
	AddJob(job *Job)
	GetANewJob() *Job
	GetJobById(jobId string) (*Job, error)
	AllocateToAgent(jobId string, agentId string) bool
	ListJobs() ([]*Job, error)
}

type JobPool struct {
	jobs  map[string]*Job
	locks map[string]string
}

func NewJobPool() *JobPool {
	jp := &JobPool{
		jobs:  make(map[string]*Job),
		locks: make(map[string]string),
	}

	return jp
}

// AddJob adds a job into the pool
func (jp *JobPool) AddJob(job *Job) {
	jp.jobs[job.Id] = job
	job.Log("Job created")
}

// GetANewJob returns a pointer to the first job with a status JOB_STATUS_NEW.
// If no jobs have such status, then nil is returned
func (jp *JobPool) GetANewJob() *Job {
	for _, job := range jp.jobs {
		if job.Status == JOB_STATUS_NEW {
			return job
		}
	}
	return nil
}

func (jp *JobPool) GetJobById(jobId string) (*Job, error) {
	job, exists := jp.jobs[jobId]
	if !exists {
		return nil, ErrJobDoesNotExist
	}
	return job, nil
}

func (jp *JobPool) LockJob(jobId string, lockId string) bool {
	_, exists := jp.locks[jobId]
	if exists {
		return false
	}
	jp.locks[jobId] = lockId
	return true
}

func (jp *JobPool) UnlockJob(jobId string, lockId string) bool {
	lock, exists := jp.locks[jobId]
	if exists {
		return false
	}

	if lock != lockId {
		return false
	}

	delete(jp.locks, jobId)
	return true
}

func (jp *JobPool) AllocateToAgent(jobId string, agentId string) bool {
	job, err := jp.GetJobById(jobId)
	if err != nil {
		return false
	}
	lockId := uuid.NewString()
	locked := jp.LockJob(job.Id, lockId)
	if !locked {
		return false
	}
	job.Status = JOB_STATUS_ALLOCATED
	job.AgentAllocated = agentId
	job.Log(fmt.Sprintf("Allocated to agent %s", agentId))
	jp.UnlockJob(job.Id, lockId)
	return true
}

func (jp *JobPool) ListJobs() ([]*Job, error) {
	list := make([]*Job, 0, len(jp.jobs))

	for _, job := range jp.jobs {
		list = append(list, job)
	}

	return list, nil
}
