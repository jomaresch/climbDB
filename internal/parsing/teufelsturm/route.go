package teufelsturm

import (
	"fmt"
	"hash/crc32"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jomaresch/climbDB/internal/model"
	"go.uber.org/zap"
)

const routeTableSelector = "body > table > tbody > tr > td > table > tbody > tr > td > div > table > tbody"
const commentTableSelector = "body > table > tbody > tr > td > table > tbody > tr > td > table > tbody"

var routeIdExtractor = regexp.MustCompile(`\d+$`)
var createdTimeExtractor = regexp.MustCompile(`\d{2}.\d{2}.\d{4} \d{2}:\d{2}$`)

func extractRouteID(regionLink string) string {
	return routeIdExtractor.FindString(regionLink)
}

func extractCreatedTime(columnText string) (time.Time, error) {
	return time.Parse("02.01.2006 15:04", createdTimeExtractor.FindString(columnText))
}

func ParseRouteList(page io.Reader, summitMapping map[string]string) ([]*model.Route, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	var routes []*model.Route

	d.Find(routeTableSelector).Children().Each(func(index int, row *goquery.Selection) {
		// Skip the table header.
		if index == 0 {
			return
		}
		// Here we parse the columns.
		route := &model.Route{}
		row.Children().Each(func(i int, col *goquery.Selection) {
			switch i {
			case 1:
				summitName := col.Children().First().Text()
				summitID, ok := summitMapping[summitName]
				if !ok {
					zap.L().Warn("no matching summit id found", zap.String("summit_name", summitName))
				}
				route.SummitID = summitID
			case 2:
				routeLinkElement := col.Children().First().Children().First()
				routeUrl, _ := routeLinkElement.Attr("href")
				route.DisplayName = routeLinkElement.Text()
				route.ID = extractRouteID(routeUrl)
			case 3:
				rating, err := ParseRatingFromPictureElement(col.Children().First().Children().First())
				if err != nil {
					zap.L().Error("failed to parse rating", zap.Error(err))
				}
				route.AvgRating = rating
			case 4:
				grading := ParseGradeFromElement(col.Children().First())
				route.GradeID = grading.GradeID
				route.GuideRating = grading.GuideRating
				route.RedPointGradeID = grading.RedPointGradeID
				route.SuggestedGradeID = grading.SuggestedGradeID
				route.JumpGrade = grading.JumpGrade
			}
		})
		routes = append(routes, route)
	})

	return routes, nil
}

func ParseRouteDetails(page io.Reader, routeID string) ([]*model.Comment, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	var comments []*model.Comment

	t := d.Find(commentTableSelector)
	t.Children().Each(func(index int, row *goquery.Selection) {
		// Skip the table header.
		if index == 0 {
			return
		}
		// Here we parse the columns.
		comment := &model.Comment{RouteID: routeID}
		row.Children().Each(func(i int, col *goquery.Selection) {
			switch i {
			case 0:
				comment.Author = col.Children().First().Text()
				createdTime, err := extractCreatedTime(col.Text())
				if err != nil {
					zap.L().Error("failed to parse created time", zap.Error(err))
				}
				comment.AuthenticatedAuthor = strings.Contains(col.Text(), "Authentifizierter")
				comment.CreatedTime = createdTime
			case 1:
				comment.Text = strings.TrimSpace(col.Text())
			case 2:
				comment.Rating = ParseRatingFromText(col.Text())
			}
		})

		commentHash := crc32.ChecksumIEEE(append(append([]byte(comment.Text), []byte(comment.RouteID)...), []byte(comment.Author)...))
		comment.ID = strconv.FormatInt(int64(commentHash), 16)
		comments = append(comments, comment)
	})
	if len(comments) == 0 {
		zap.L().Error("no comments for route", zap.String("route", routeID))
	}
	return comments, nil
}
