package model

import (
	"time"
)

type TeufelsturmRating int
type GuideRating int

type Comment struct {
	ID                  string
	Author              string
	CreatedTime         time.Time
	AuthenticatedAuthor bool
	Rating              TeufelsturmRating
	Text                string
}

type Route struct {
	ID             string
	SummitID       string
	DisplayName    string
	AvgRating      TeufelsturmRating
	GuideRating    GuideRating
	Grade          Grade
	SuggestedGrade Grade
	RedPointGrade  Grade
}

type Summit struct {
	ID          string
	GuideID     string
	RegionID    string
	DisplayName string
	Lat         float64
	Long        float64
}

type Region struct {
	ID          string
	DisplayName string
}
