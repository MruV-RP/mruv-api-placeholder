package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/jobs"
)

type JobsServer struct {
	gen generator.IGenerator
}

func NewJobsServer(gen generator.IGenerator) *JobsServer {
	return &JobsServer{gen: gen}
}

func (j *JobsServer) CreateJob(ctx context.Context, request *jobs.CreateJobRequest) (*jobs.CreateJobResponse, error) {
	panic("implement me")
}

func (j *JobsServer) GetJob(ctx context.Context, request *jobs.GetJobRequest) (*jobs.GetJobResponse, error) {
	panic("implement me")
}

func (j *JobsServer) UpdateJob(ctx context.Context, request *jobs.UpdateJobRequest) (*jobs.UpdateJobResponse, error) {
	panic("implement me")
}

func (j *JobsServer) DeleteJob(ctx context.Context, request *jobs.DeleteJobRequest) (*jobs.DeleteJobResponse, error) {
	panic("implement me")
}
