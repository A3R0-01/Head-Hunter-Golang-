package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minCompanyName = 5
	minAddress     = 10
)

type Company struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Email        []string           `bson:"email" json:"email"`
	Address      string             `bson:"address" json:"address"`
	Telephone    []string           `bson:"telephone" json:"telephone"`
	HeadOfHr     primitive.ObjectID `bson:"headOfHr" json:"headOfHr"`
	Verified     bool               `bson:"verified" json:"verified"`
	Verification primitive.ObjectID `bson:"verification" json:"verification"`
	Created      time.Time          `bson:"Created" json:"Created"`
}
type CreateCompanyParams struct {
	Name      string           `bson:"name" json:"name"`
	Email     []string         `bson:"email" json:"email"`
	Address   string           `bson:"address" json:"address"`
	Telephone []string         `bson:"telephone" json:"telephone"`
	HeadOfHr  CreateUserParams `bson:"headOfHr" json:"headOfHr"`
}
type UpdateCompanyParams struct {
	Name      string   `bson:"name" json:"name"`
	Email     []string `bson:"Email" json:"Email"`
	Address   string   `bson:"address" json:"address"`
	Telephone []string `bson:"telephone" json:"telephone"`
}

func (c *UpdateCompanyParams) ToMongoBson() (bson.M, error) {
	updateParams := bson.M{}
	emailList := []string{}
	for _, email := range c.Email {
		if IsEmailValid(email) {
			updateParams["email"] = append(emailList, email)
		} else if email != "" {
			return nil, fmt.Errorf("invalid email: %v", email)
		}
	}
	if !(len(c.Name) < minCompanyName) {
		updateParams["name"] = c.Name
	}
	if !(len(c.Address) < minAddress) {
		updateParams["address"] = c.Address
	}
	phoneList := []string{}
	for _, phone := range c.Telephone {
		if IsPhoneValid(phone) {
			updateParams["phone"] = append(phoneList, phone)
		} else if phone != "" {
			return nil, fmt.Errorf("invalid number: %v", phone)
		}
	}
	return updateParams, nil
}

func (c *CreateCompanyParams) Validate() map[string]any {
	errors := map[string]any{}

	for _, email := range c.Email {
		if !IsEmailValid(email) {
			errors["email"] = fmt.Sprint("invalid email: ", email)
		}
	}
	if len(c.Name) < minCompanyName {
		errors["name"] = "short name"
	}
	if len(c.Address) < minAddress {
		errors["address"] = "invalid address"
	}
	for _, phone := range c.Telephone {
		if !IsPhoneValid(phone) {
			errors["phone"] = fmt.Sprint("phone number is not valid: ", phone)
		}
	}
	if len(c.HeadOfHr.Validate()) > 0 {
		errors["headOfHr"] = c.HeadOfHr.Validate()
	}

	return errors
}
func (c *CreateCompanyParams) FromParams() *Company {
	return &Company{
		Name:      c.Name,
		Email:     c.Email,
		Address:   c.Address,
		Telephone: c.Telephone,
		Verified:  false,
		Created:   time.Now(),
	}
}
