package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type SessionStore interface {
	GetSessionByID(context.Context, string) (*types.Session, error)
	GetSessions(context.Context, Map) ([]*types.Session, error)
	CreateSession(context.Context, *types.Session) (*types.Session, error)
	UpdateSession(context.Context, string) error
	DeleteSessions(context.Context, Map) error
}
