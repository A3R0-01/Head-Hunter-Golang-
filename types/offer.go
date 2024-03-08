package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const descriptionLen = 20

type Offer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	Recruiter    primitive.ObjectID `bson:"Recruiter" json:"Recruiter"`
	JobSeeker    primitive.ObjectID `bson:"JobSeeker" json:"JobSeeker"`
	Title        string             `bson:"Title" json:"Title"`
	Introduction string             `bson:"Introduction" json:"Introduction"`
	Description  []string           `bson:"Description" json:"Description"`
	Salary       string             `bson:"Salary" json:"Salary"`
	Status       bool               `bson:"Status" json:"Status"`
	DueDate      time.Time          `bson:"DueDate" json:"DueDate"`
	Created      time.Time          `bson:"Created" json:"Created"`
}

type CreateOfferParams struct {
	Recruiter    string    `json:"Recruiter"`
	JobSeeker    string    `json:"JobSeeker"`
	Title        string    `json:"Title"`
	Introduction string    `json:"Introduction"`
	Description  []string  `json:"Description"`
	Salary       string    `json:"Salary"`
	DueDate      time.Time `json:"DueDate"`
}

func (c *CreateOfferParams) FromParams() (*Offer, error) {
	oidRecruiter, err := primitive.ObjectIDFromHex(c.Recruiter)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}
	oidJobSeeker, err := primitive.ObjectIDFromHex(c.JobSeeker)
	if err != nil {
		return nil, fmt.Errorf("invalid job seeker")
	}
	if len(c.Title) < minTitle {
		return nil, fmt.Errorf("small title")
	}
	if len(c.Description) > 0 {
		for _, d := range c.Description {
			if len(d) < descriptionLen {
				return nil, fmt.Errorf("small description, requires %v", descriptionLen)
			}
		}
	}
	if len(c.Introduction) < minIntroduction {
		return nil, fmt.Errorf("introduction too small")
	}

	return &Offer{
		Recruiter:    oidRecruiter,
		JobSeeker:    oidJobSeeker,
		Title:        c.Title,
		Introduction: c.Introduction,
		Description:  c.Description,
		Salary:       c.Salary,
		Status:       false,
		DueDate:      c.DueDate,
	}, nil
}
