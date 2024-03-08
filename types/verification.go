package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultVerificationYears  = 5
	DefaultVerificationMonths = 5
	DefaultVerificationDays   = 5
)

type Verification struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"Employee,omitempty" json:"Employee,omitempty"`
	CompanyID  primitive.ObjectID `bson:"Company,omitempty" json:"Company,omitempty"`
	Document   primitive.ObjectID `bson:"Document" json:"Document"`
	ValidTill  time.Time          `bson:"ValidTill" json:"ValidTill"`
	Created    time.Time          `bson:"Created,omitempty" json:"Created,omitempty"`
}

type CreateVerificationParams struct {
	EmployeeID string `json:"Employee,omitempty"`
	CompanyID  string `json:"Company,omitempty"`
	Document   string `json:"Document"`
}

func (c *CreateVerificationParams) FromParams() (*Verification, error) {
	oidEmployeeID, err := primitive.ObjectIDFromHex(c.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}
	oidCompanyID, err := primitive.ObjectIDFromHex(c.CompanyID)
	if err != nil {
		return nil, fmt.Errorf("invalid company")
	}
	oidDocument, err := primitive.ObjectIDFromHex(c.Document)
	if err != nil {
		return nil, fmt.Errorf("invalid file")
	}
	return &Verification{
		EmployeeID: oidEmployeeID,
		CompanyID:  oidCompanyID,
		Document:   oidDocument,
		ValidTill:  time.Now().AddDate(DefaultVerificationYears, DefaultVerificationMonths, DefaultVerificationDays),
		Created:    time.Now(),
	}, nil
}
