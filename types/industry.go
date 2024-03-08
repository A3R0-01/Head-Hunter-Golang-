package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minIndustryName = 5
)

type Industry struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"Industry" json:"Industry"`
	Users   int                `bson:"Users" json:"Users"`
	Created time.Time          `bson:"Created" json:"Created"`
}

type CreateIndustryParams struct {
	Name string `json:"Name"`
}

func (c *CreateIndustryParams) FromParams() (*Industry, error) {
	if len(c.Name) < minIndustryName {
		return nil, fmt.Errorf("name is too short")
	}
	return &Industry{
		Name:    c.Name,
		Users:   0,
		Created: time.Now(),
	}, nil
}
