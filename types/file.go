package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	JobSeeker primitive.ObjectID `bson:"JobSeeker" json:"JobSeeker"`
	File      string             `bson:"File" json:"File"`
	Confirm   bool               `bson:"Confirm" json:"Confirm"`
	Created   time.Time          `bson:"Created" json:"Created"`
}

type CreateFileParams struct {
	JobSeeker string `json:"JobSeeker"`
	File      string `json:"File"`
	Confirm   bool   `json:"Confirm"`
}

func (c *CreateFileParams) FromParams() (*File, error) {
	oidJobSeeker, err := primitive.ObjectIDFromHex(c.JobSeeker)
	if err != nil {
		return nil, fmt.Errorf("invalid job seeker")
	}
	return &File{
		JobSeeker: oidJobSeeker,
		File:      c.File,
		Confirm:   false,
		Created:   time.Now(),
	}, nil
}
