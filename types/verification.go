package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultVerificationYears  = 5
	DefaultVerificationMonths = 5
	DefaultVerificationDays   = 5
)

type Verification struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"EmployeeID,omitempty" json:"EmployeeID,omitempty"`
	CompanyID  primitive.ObjectID `bson:"CompanyID,omitempty" json:"CompanyID,omitempty"`
	Document   primitive.ObjectID `bson:"DocumentID" json:"DocumentID"`
	ValidTill  time.Time          `bson:"ValidTill" json:"ValidTill"`
	Created    time.Time          `bson:"Created,omitempty" json:"Created,omitempty"`
}

type CreateVerificationParams struct {
	EmployeeID string `json:"EmployeeID,omitempty"`
	CompanyID  string `json:"CompanyID,omitempty"`
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

type UpdateVerificationParams struct {
	EmployeeID string    `bson:"EmployeeID,omitempty" json:"EmployeeID,omitempty"`
	Document   string    `bson:"DocumentID" json:"DocumentID"`
	ValidTill  time.Time `bson:"ValidTill" json:"ValidTill"`
}

func (u *UpdateVerificationParams) Validate() error {
	return nil
}

func (u *UpdateVerificationParams) ToUpdateMongo() (bson.M, []error) {
	update := bson.M{}
	errors := []error{}
	_, err := primitive.ObjectIDFromHex(u.EmployeeID)
	if err == nil {
		update["EmployeeID"] = u.EmployeeID
	} else {
		errors = append(errors, err)
	}
	_, err = primitive.ObjectIDFromHex(u.Document)
	if err == nil {
		update["Document"] = u.Document
	} else {
		errors = append(errors, err)
	}
	if u.ValidTill.After(time.Now()) {
		update["ValidTill"] = u.ValidTill
	} else {
		errors = append(errors, err)
	}
	return update, errors
}
