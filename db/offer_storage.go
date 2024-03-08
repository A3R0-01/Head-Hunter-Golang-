package db

import (
	"context"

	"github.com/A3R0-01/head-hunter/types"
)

type OfferStore interface {
	GetOfferByID(context.Context, string) (*types.Offer, error)
	GetOffers(context.Context, Map) ([]*types.Offer, error)
	CreateOffer(context.Context, *types.Offer) (*types.Offer, error)
	DeleteOffer(context.Context, string) error
	UpdateOffer(context.Context, string, Map) error
}
