package teufelsturm

import (
	"fmt"
	"io"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/jomaresch/climbDB/internal/model"
)

const regionOverviewTableSelector = "body > table > tbody > tr > td > table > tbody > tr > td > table > tbody"

var regionIdExtractor = regexp.MustCompile(`\d+$`)

func extractRegionID(regionLink string) string {
	return regionIdExtractor.FindString(regionLink)
}

func ParseRegionList(page io.Reader) ([]*model.Region, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	var summits []*model.Region
	d.Find(regionOverviewTableSelector).Children().Each(func(index int, row *goquery.Selection) {
		// Skip the table header.
		if index == 0 {
			return
		}

		domLinkElement := row.Find("td > font > a")
		regionUrl, _ := domLinkElement.Attr("href")

		regions := &model.Region{
			ID:          extractRegionID(regionUrl),
			DisplayName: domLinkElement.Text(),
		}

		summits = append(summits, regions)
	})

	return summits, nil
}
