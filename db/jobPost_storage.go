package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type JobPostStore interface {
	Dropper
	GetJobPostByID(context.Context, string) (*types.JobPost, error)
	GetJobPosts(context.Context, Map) ([]*types.JobPost, error)
	CreateJobPost(context.Context, *types.JobPost) (*types.JobPost, error)
	DeleteJobPost(context.Context, string) error
	UpdateJobPost(context.Context, string, types.UpdateJobPostParams) error
}
