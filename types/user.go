package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 8
	minFirstNameLen = 3
	minLastNameLen  = 3
	minPasswordLen  = 8
	minNatIDLen     = 12
	minLocationLen  = 8
	maxNatIDLen     = 20
	minPhoneLen     = 7
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	Company           primitive.ObjectID `bson:"Company" json:"Company"`
	FirstName         string             `bson:"FirstName" json:"FirstName"`
	LastName          string             `bson:"LastName" json:"LastName"`
	Email             string             `bson:"Email" json:"Email"`
	NatID             string             `bson:"NatID" json:"NatID"`
	DateOfBirth       time.Time          `bson:"DateOfBirth" json:"DateOfBirth"`
	Location          string             `bson:"Location" json:"Location"`
	Phone             string             `bson:"Phone" json:"Phone"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
	Industry          primitive.ObjectID `bson:"Industry" json:"Industry"`
	Student           bool               `bson:"Student" json:"Student"`
	JobSeeker         bool               `bson:"JobSeeker" json:"JobSeeker"`
	Recruiter         bool               `bson:"Recruiter" json:"Recruiter"`
	Admin             bool               `bson:"Admin" json:"_"`
	Created           time.Time          `bson:"Created" json:"Created"`
}
type CreateUserParams struct {
	FirstName   string    ` json:"FirstName"`
	LastName    string    `json:"LastName"`
	Email       string    `json:"Email"`
	NatID       string    `json:"NatID"`
	DateOfBirth time.Time `json:"DateOfBirth"`
	Location    string    `json:"Location"`
	Phone       string    `json:"Phone"`
	Password    string    `json:"Password"`
	Industry    string    `json:"Industry"`
	Student     bool      `json:"Student"`
	JobSeeker   bool      `json:"JobSeeker"`
	Recruiter   bool      `json:"Recruiter"`
}

func (c *CreateUserParams) FromParams() (*User, error) {
	encpsw, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	oid, err := primitive.ObjectIDFromHex(c.Industry)
	if err != nil {
		return nil, fmt.Errorf("invalid industry")
	}
	return &User{
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Email:             c.Email,
		NatID:             c.NatID,
		DateOfBirth:       c.DateOfBirth,
		Location:          c.Location,
		Phone:             c.Phone,
		EncryptedPassword: string(encpsw),
		Industry:          oid,
		Student:           c.Student,
		JobSeeker:         c.JobSeeker,
		Recruiter:         c.Recruiter,
		Created:           time.Now(),
	}, nil
}
func (c *CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	now := time.Now()
	if len(c.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("first name should be at least %d characters", minFirstNameLen)
	}
	if len(c.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("last name should be at least %d characters", minLastNameLen)
	}
	if len(c.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password should be at least %d characters", minPasswordLen)
	}
	if !IsEmailValid(c.Email) {
		errors["email"] = "invalid Email"
	}
	if len(c.NatID) < minNatIDLen || len(c.NatID) > maxNatIDLen {
		errors["NatID"] = "Invalid National ID number"
	}
	if c.DateOfBirth.After(now) || c.DateOfBirth.After(now.AddDate(-18, 0, 0)) {
		errors["DateOfBirth"] = "Invalid date of birth or too young"
	}
	if len(c.Location) < minLocationLen {
		errors["Location"] = "short location"
	}
	if !IsPhoneValid(c.Phone) {
		errors["Phone"] = "invalid phone number"
	}
	_, err := primitive.ObjectIDFromHex(c.Industry)
	if err != nil {
		errors["industry"] = "invalid industry"
	}
	return errors
}

type UpdateUserParams struct {
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Email       string    `json:"Email"`
	NatID       string    `json:"NatID"`
	DateOfBirth time.Time `json:"DateOfBirth"`
	Location    string    `json:"Location"`
	Phone       string    `json:"Phone"`
	Password    string    `json:"Password"`
	Industry    string    `json:"Industry"`
	Student     bool      `json:"Student"`
	JobSeeker   bool      `json:"JobSeeker"`
	Recruiter   bool      `json:"Recruiter"`
}

func (c *UpdateUserParams) ToUpdateMongo() (bson.M, error) {
	updateParams := bson.M{}
	now := time.Now()
	if !(len(c.FirstName) < minFirstNameLen) {
		updateParams["FirstName"] = c.FirstName
	}
	if !(len(c.LastName) < minLastNameLen) {
		updateParams["LastName"] = c.LastName
	}
	if !(len(c.Password) < minPasswordLen) {
		encpsw, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcryptCost)
		if err != nil {
			return nil, err
		}
		updateParams["Password"] = encpsw
	}
	if IsEmailValid(c.Email) {
		updateParams["Email"] = c.Email
	}
	if !(len(c.NatID) < minNatIDLen || len(c.NatID) > maxNatIDLen) {
		updateParams["NatID"] = c.NatID
	}
	if !(c.DateOfBirth.After(now) || c.DateOfBirth.After(now.AddDate(-18, 0, 0))) {
		updateParams["DateOfBirth"] = c.DateOfBirth
	}
	if len(c.Location) > minLocationLen {
		updateParams["Location"] = c.Location
	}
	if IsPhoneValid(c.Phone) {
		updateParams["Phone"] = c.Phone
	}
	oid, err := primitive.ObjectIDFromHex(c.Industry)
	if err != nil {
		updateParams["Industry"] = oid
	}
	return updateParams, nil
}
