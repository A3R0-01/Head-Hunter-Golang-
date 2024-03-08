package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DutiesTextLen         = 6
	minDuties             = 1
	minTitle              = 6
	minIntroduction       = 20
	minQualifications     = 3
	minQualificationsText = 8
	minApplicants         = 10
	minHowToApply         = 10
)

type JobPost struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	Recruiter         primitive.ObjectID `bson:"Recruiter" json:"Recruiter"`
	Title             string             `bson:"Title" json:"Title"`
	Location          string             `bson:"Location" json:"Location"`
	Introduction      string             `bson:"Introduction" json:"Introduction"`
	Duties            []string           `bson:"Duties" json:"Duties"`
	Qualifications    []string           `bson:"Education" json:"Education"`
	Notes             []string           `bson:"Notes" json:"Notes"`
	EasyApply         bool               `bson:"EasyApply" json:"EasyApply"`
	HowToApply        string             `bson:"HowToApply" json:"HowToApply"`
	MaximumApplicants int                `bson:"MaximumApplicants" json:"MaximumApplicants"`
	Industry          primitive.ObjectID `bson:"Industry" json:"Industry"`
	DueDate           time.Time          `bson:"DueDate" json:"DueDate"`
	Created           time.Time          `bson:"Created" json:"Created"`
}

type CreateJobPostParams struct {
	Recruiter         string    `bson:"Recruiter" json:"Recruiter"`
	Title             string    `bson:"Title" json:"Title"`
	Location          string    `bson:"Location" json:"Location"`
	Introduction      string    `bson:"Introduction" json:"Introduction"`
	Duties            []string  `bson:"Duties" json:"Duties"`
	Qualifications    []string  `bson:"Education" json:"Education"`
	Notes             []string  `bson:"Notes" json:"Notes"`
	EasyApply         bool      `bson:"EasyApply" json:"EasyApply"`
	HowToApply        string    `bson:"HowToApply" json:"HowToApply"`
	MaximumApplicants int       `bson:"MaximumApplicants" json:"MaximumApplicants"`
	Industry          string    `bson:"Industry" json:"Industry"`
	DueDate           time.Time `bson:"DueDate" json:"DueDate"`
}
type UpdateJobPostParams struct {
	Title             string    `bson:"Title" json:"Title"`
	Location          string    `bson:"Location" json:"Location"`
	Introduction      string    `bson:"Introduction" json:"Introduction"`
	Duties            []string  `bson:"Duties" json:"Duties"`
	Qualifications    []string  `bson:"Education" json:"Education"`
	Notes             []string  `bson:"Notes" json:"Notes"`
	EasyApply         bool      `bson:"EasyApply" json:"EasyApply"`
	HowToApply        string    `bson:"HowToApply" json:"HowToApply"`
	MaximumApplicants int       `bson:"MaximumApplicants" json:"MaximumApplicants"`
	Industry          string    `bson:"Industry" json:"Industry"`
	DueDate           time.Time `bson:"DueDate" json:"DueDate"`
}

func (u *UpdateJobPostParams) ToMongoBson() (bson.M, error) {
	params := bson.M{}
	if len(u.Duties) > minDuties {
		dutiesList := []string{}
		for _, duty := range u.Duties {
			if !(len(duty) < DutiesTextLen) {
				params["Duties"] = append(dutiesList, duty)
			}
		}
	} else {
		return nil, fmt.Errorf("add more duties")
	}
	if len(u.Qualifications) > minQualifications {
		for _, qualification := range u.Qualifications {
			qualificationsList := []string{}
			if len(qualification) < minQualificationsText {
				params["Qualifications"] = append(qualificationsList, qualification)
			}
		}
	} else {
		return nil, fmt.Errorf("add more qualifications")
	}

	if !(len(u.Title) < minTitle) {
		params["Title"] = u.Title
	}
	if len(u.Introduction) < minIntroduction {
		params["Introduction"] = u.Introduction
	}
	if len(u.Location) < minLocationLen {
		params["Location"] = u.Location
	}
	if !(u.DueDate.Before(time.Now())) {
		params["DueDate"] = u.DueDate
	}
	oidIndustry, err := primitive.ObjectIDFromHex(u.Industry)
	if err == nil {
		params["Industry"] = oidIndustry
	}
	if !(len(u.HowToApply) < minHowToApply) {
		params["HowToApply"] = u.HowToApply
	}
	if !(u.MaximumApplicants < minApplicants) {
		params["MaximumApplicants"] = u.MaximumApplicants
	}

	return params, nil
}

func (j *CreateJobPostParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(j.Duties) > minDuties {
		for _, duty := range j.Duties {
			if len(duty) < DutiesTextLen {
				errors["dutiesText"] = fmt.Sprint("invalid duty: ", duty)
			}
		}
	} else {
		errors["duties"] = "insufficient duties"
	}
	if len(j.Qualifications) > minQualifications {
		for _, qualification := range j.Qualifications {
			if len(qualification) < minQualificationsText {
				errors["qualificationsText"] = fmt.Sprint("short qualification: ", qualification)
			}
		}
	} else {
		errors["qualifications"] = "insufficient qualifications"
	}

	if len(j.Title) < minTitle {
		errors["title"] = "short title"
	}
	if len(j.Introduction) < minIntroduction {
		errors["introduction"] = "invalid introduction"
	}
	if len(j.Location) < minLocationLen {
		errors["location"] = "invalid location"
	}
	if j.DueDate.Before(time.Now()) {
		errors["dueDate"] = "invalid time"
	}
	_, err := primitive.ObjectIDFromHex(j.Industry)
	if err != nil {
		errors["industry"] = "invalid industry"
	}
	_, err = primitive.ObjectIDFromHex(j.Recruiter)
	if err != nil {
		errors["industry"] = "invalid recruiter"
	}
	if len(j.HowToApply) < 1 {
		j.EasyApply = true
	} else if len(j.HowToApply) < minHowToApply {
		errors["howToApply"] = "add more info"
	}
	if j.MaximumApplicants < minApplicants {
		errors["maximumApplicants"] = "applicants should be more than 10"
	}
	return errors
}

func (j *CreateJobPostParams) FromParams() (*JobPost, error) {
	oidIndustry, err := primitive.ObjectIDFromHex(j.Industry)
	if err != nil {
		return nil, fmt.Errorf("Invalid Industry")
	}
	oidRecruiter, err := primitive.ObjectIDFromHex(j.Recruiter)
	if err != nil {
		return nil, fmt.Errorf("invalid recruiter")
	}

	return &JobPost{
		Recruiter:         oidRecruiter,
		Title:             j.Title,
		Location:          j.Location,
		Introduction:      j.Introduction,
		Duties:            j.Duties,
		Qualifications:    j.Qualifications,
		Notes:             j.Notes,
		EasyApply:         j.EasyApply,
		HowToApply:        j.HowToApply,
		MaximumApplicants: j.MaximumApplicants,
		Industry:          oidIndustry,
		DueDate:           j.DueDate,
		Created:           time.Now(),
	}, nil
}
