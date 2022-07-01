package teufelsturm

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jomaresch/climbDB/internal/model"
	"go.uber.org/zap"
)

type TeufelsturmGrading struct {
	GradeID          int
	SuggestedGradeID int
	RedPointGradeID  int
	GuideRating      int
	JumpGrade        int
}

func ParseGradeFromElement(element *goquery.Selection) *TeufelsturmGrading {
	grading := &TeufelsturmGrading{}
	gradingText := strings.TrimSpace(element.Text())
	gradingElements := strings.Split(gradingText, " ")

	nextGradeIsRedPoint := false
	for _, g := range gradingElements {
		if strings.ContainsAny(g, "!*") {
			switch g {
			case "!":
				grading.GuideRating = -1
			case "!!":
				grading.GuideRating = -2
			case "*":
				grading.GuideRating = 1
			case "**":
				grading.GuideRating = 2
			}
			continue
		}

		if strings.Contains(g, "RP") {
			nextGradeIsRedPoint = true
			continue
		}

		if nextGradeIsRedPoint {
			grading.RedPointGradeID = getGradeIDForSaxonyGrade(g)
			nextGradeIsRedPoint = false
			continue
		}

		if strings.ContainsAny(g, "()") {
			cleanGrade := strings.ReplaceAll(g, "(", "")
			cleanGrade = strings.ReplaceAll(cleanGrade, ")", "")
			grading.SuggestedGradeID = getGradeIDForSaxonyGrade(cleanGrade)
			continue
		}

		if strings.Contains(g, "/") {
			jumpGradeList := strings.Split(g, "/")
			jumpGrade, err := strconv.Atoi(jumpGradeList[0])
			if err != nil {
				zap.L().Error("failed to parse jump grade", zap.Error(err), zap.Any("grade", g))
			}
			grading.JumpGrade = jumpGrade
			g = jumpGradeList[1]
			grading.GradeID = getGradeIDForSaxonyGrade(g)
			continue
		}

		if jumpGrade, err := strconv.Atoi(g); err == nil {
			grading.JumpGrade = jumpGrade
			continue
		}

		if strings.HasPrefix(strings.ToLower(g), "s") {
			jumpGrade, _ := strconv.Atoi(strings.TrimPrefix(strings.ToLower(g), "s"))
			grading.JumpGrade = jumpGrade
			continue
		}

		if strings.ContainsAny(g, "-?") {
			continue
		}

		grading.GradeID = getGradeIDForSaxonyGrade(g)
	}
	return grading
}

func getGradeIDForSaxonyGrade(saxonyGrade string) int {
	grade, found := model.GradesSaxony[saxonyGrade]
	if !found {
		zap.L().Error("grade not found in grade list", zap.String("grade", saxonyGrade))
		return 0
	}
	return grade.ID
}
