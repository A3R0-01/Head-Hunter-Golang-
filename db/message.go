package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type MessageStore interface {
	GetMessageByID(context.Context, string) (*types.Message, error)
	GetMessages(context.Context, Map) ([]*types.Message, error)
	CreateMessage(context.Context, *types.Message) (*types.Message, error)
	DeleteMessages(context.Context, Map) error
}
