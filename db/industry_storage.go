package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type IndustryStore interface {
	GetIndustryByID(context.Context, string) (*types.Industry, error)
	GetIndustries(context.Context, Map) ([]*types.Industry, error)
	CreateIndustry(context.Context, *types.Industry) (*types.Industry, error)
	DeleteIndustry(context.Context, string) error
	UpdateIndustry(context.Context, string, types.UpdateUserParams) error
}
