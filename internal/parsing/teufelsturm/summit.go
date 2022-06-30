package teufelsturm

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jomaresch/climbDB/internal/model"
	"go.uber.org/zap"
)

const summitOverviewTableSelector = "body > table > tbody > tr > td > table > tbody > tr > td > div > table > tbody"

func ParseSummitList(page io.Reader, areaMap map[string]string) ([]*model.Summit, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	var summits []*model.Summit

	d.Find(summitOverviewTableSelector).Children().Each(func(index int, row *goquery.Selection) {
		// Skip the table header.
		if index == 0 {
			return
		}
		// Here we parse the columns.
		summit := &model.Summit{}
		row.Children().Each(func(i int, col *goquery.Selection) {
			switch i {
			case 0:
				summit.ID = col.Text()
			case 1:
				summit.DisplayName = col.Text()
			case 2:
				summit.GuideID = col.Text()
			case 3:
				regionID, _ := areaMap[col.Text()]
				summit.RegionID = regionID
			}
		})
		summits = append(summits, summit)
	})

	return summits, nil
}

const summitDetailsSelector = "body > table > tbody > tr > td > font > table > tbody"

func ParseSummitDetails(page io.Reader) (*model.Summit, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	summit := &model.Summit{}

	d.Find(summitDetailsSelector).Children().Each(func(index int, row *goquery.Selection) {
		previousElementText := ""
		row.Children().Each(func(i int, col *goquery.Selection) {
			switch previousElementText {
			case "Longitude":
				long, err := strconv.ParseFloat(col.Text(), 10)
				if err != nil {
					zap.L().Error("failed to parse longitude", zap.Error(err))
				}
				summit.Long = long
			case "Latitude":
				lat, err := strconv.ParseFloat(col.Text(), 10)
				if err != nil {
					zap.L().Error("failed to parse longitude", zap.Error(err))
				}
				summit.Lat = lat
			}
			previousElementText = strings.TrimSpace(col.Text())
		})
	})

	return summit, nil
}
