package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type ApplicationStore interface {
	GetApplicationByID(context.Context, string) (*types.Application, error)
	GetApplications(context.Context, Map) ([]*types.Application, error)
	CreateApplication(context.Context, *types.Application) (*types.Application, error)
	AlterApplication(context.Context) (*types.Application, error)
	DeleteApplication(context.Context, string) error
}
