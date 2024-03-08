package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type FileStore interface {
	GetFileByID(context.Context, string) (*types.File, error)
	GetFiles(context.Context, Map) ([]*types.File, error)
	CreateFile(context.Context, *types.File) (*types.File, error)
	DeleteFile(context.Context, string) error
	UpdateFile(context.Context, string, Map) error
}
