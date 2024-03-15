package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minIndustryName = 5
)

type Industry struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"Industry" json:"Industry"`
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
		Created: time.Now(),
	}, nil
}

type UpdateIndustryParams struct {
	Name  string `bson:"Industry" json:"Industry"`
	Users int    `bson:"Users" json:"Users"`
}

func (u *UpdateIndustryParams) ToMongoBson() {
	values := bson.M{}
	if len(u.Name) > minIndustryName {
		values["Name"] = u.Name
	}
}
