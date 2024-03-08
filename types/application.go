package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"ID,omitempty"`
	JobSeeker   primitive.ObjectID   `bson:"JobSeeker" json:"JobSeeker"`
	JobPost     primitive.ObjectID   `bson:"JobPost" json:"JobPost"`
	Qualified   bool                 `bson:"Qualified" json:"Qualified"`
	Failed      bool                 `bson:"Failed" json:"Failed"`
	Resume      primitive.ObjectID   `bson:"Resume" json:"Resume"`
	CoverLetter primitive.ObjectID   `bson:"CoverLetter" json:"CoverLetter"`
	Additional  []primitive.ObjectID `bson:"Additional" json:"Additional"`
	Created     time.Time            `bson:"Created" json:"Created"`
}

type CreateApplicationParams struct {
	JobSeeker   string   `bson:"JobSeeker" json:"JobSeeker"`
	JobPost     string   `bson:"JobPost" json:"JobPost"`
	Qualified   bool     `bson:"Qualified" json:"Qualified"`
	Failed      bool     `bson:"Failed" json:"Failed"`
	Resume      string   `bson:"Resume" json:"Resume"`
	CoverLetter string   `bson:"CoverLetter" json:"CoverLetter"`
	Additional  []string `bson:"Additional" json:"Additional"`
}

func (c *CreateApplicationParams) FromParams() (*Application, error) {
	oidJobPost, err := primitive.ObjectIDFromHex(c.JobPost)
	if err != nil {
		return nil, fmt.Errorf("invalid job post")
	}
	oidJobSeeker, err := primitive.ObjectIDFromHex(c.JobSeeker)
	if err != nil {
		return nil, fmt.Errorf("invalid job seeker")
	}
	oidResume, err := primitive.ObjectIDFromHex(c.Resume)
	if err != nil {
		return nil, fmt.Errorf("invalid resume")
	}
	oidCoverLetter, err := primitive.ObjectIDFromHex(c.CoverLetter)
	if err != nil {
		return nil, fmt.Errorf("invalid cover Letter")
	}
	oidAdditionalList := []primitive.ObjectID{}
	if len(c.Additional) > 0 {
		for _, item := range c.Additional {
			oidAdditionalItem, err := primitive.ObjectIDFromHex(item)
			if err != nil {
				return nil, fmt.Errorf("invalid file")
			}
			oidAdditionalList = append(oidAdditionalList, oidAdditionalItem)
		}
	}

	return &Application{
		JobSeeker:   oidJobSeeker,
		JobPost:     oidJobPost,
		Qualified:   c.Qualified,
		Failed:      c.Failed,
		Resume:      oidResume,
		CoverLetter: oidCoverLetter,
		Additional:  oidAdditionalList,
		Created:     time.Now(),
	}, nil
}
