package model

import (
	"fmt"
	"time"
)

type Comment struct {
	ID                  string
	Author              string
	CreatedTime         time.Time
	AuthenticatedAuthor bool
	Rating              int
	Text                string
}

type Route struct {
	ID               string
	SummitID         string
	DisplayName      string
	AvgRating        int
	GuideRating      int
	GradeID          int
	SuggestedGradeID int
	RedPointGradeID  int
	JumpGrade        int
}

type Summit struct {
	ID          string
	GuideID     string
	RegionID    string
	DisplayName string
	Lat         float64
	Long        float64
}

func (s *Summit) String() string {
	return fmt.Sprintf("%s, %s, %s \n", s.GuideID, s.DisplayName, s.RegionID)
}

type Region struct {
	ID          string
	DisplayName string
}
