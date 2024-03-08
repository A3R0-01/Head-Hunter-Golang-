package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type CompanyStore interface {
	Dropper
	GetCompanyByID(context.Context, string) (*types.Company, error)
	GetCompanies(context.Context, Map) ([]*types.Company, error)
	CreateCompany(context.Context, *types.Company) (*types.Company, error)
	CreateRecruiterToken(context.Context, *types.Company) (string, error)
	GetRecruiters(context.Context, string) ([]*types.User, error)
	DeleteCompany(context.Context, string) error
	UpdateCompany(context.Context, string, types.UpdateCompanyParams) error
}
