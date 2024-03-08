package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type VerificationStore interface {
	GetVerificationByID(context.Context, string) (*types.Verification, error)
	GetVerifications(context.Context, Map) ([]*types.Verification, error)
	CreateVerification(context.Context, *types.Verification) (*types.Verification, error)
	DeleteVerification(context.Context, string) error
	// UpdateVerification(context.Context, string, types.UpdateVerificaParams) error
}
