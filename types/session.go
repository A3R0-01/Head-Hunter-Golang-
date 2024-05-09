package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	Recruiter primitive.ObjectID `bson:"Recruiter" json:"Recruiter"`
	JobSeeker primitive.ObjectID `bson:"JobSeeker" json:"JobSeeker"`
	Open      bool               `bson:"Open" json:"Open"`
	Created   time.Time          `bson:"Created" json:"Created"`
}
type CreateSessionParams struct {
	Recruiter string `json:"Recruiter"`
	JobSeeker string `json:"JobSeeker"`
	Open      bool   `json:"Open"`
}

func (c *CreateSessionParams) FromParams() (*Session, error) {
	oidRecruiter, err := primitive.ObjectIDFromHex(c.Recruiter)
	if err != nil {
		return nil, fmt.Errorf("invalid session")

	}
	oidJobSeeker, err := primitive.ObjectIDFromHex(c.JobSeeker)
	if err != nil {
		return nil, fmt.Errorf("invalid jobseeker")
	}
	return &Session{
		Recruiter: oidRecruiter,
		JobSeeker: oidJobSeeker,
		Open:      true,
		Created:   time.Now(),
	}, nil
}

type UpdateSessionParams struct {
	Open bool `json:"Open"`
}

func (u *UpdateSessionParams) ToUpdateMongo() bson.M {
	return bson.M{
		"Open": u.Open,
	}
}
