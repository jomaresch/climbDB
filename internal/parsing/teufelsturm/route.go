package teufelsturm

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/jomaresch/climbDB/internal/model"
)

const SummitOverviewTable = "body > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr > td > div > table > tbody"

func ParseSummitList(ctx context.Context, page io.Reader, areaMap map[string]string) ([]*model.Summit, error) {
	d, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create document from reader: %w", err)
	}
	var parsingErrors []error
	var summits []*model.Summit

	d.Find(SummitOverviewTable).Children().Each(func(index int, row *goquery.Selection) {
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
				regionID, ok := areaMap[col.Text()]
				if !ok {
					parsingErrors = append(parsingErrors, fmt.Errorf("no Id for region '%s' found", col.Text()))
					return
				}
				summit.Region = regionID
			}
		})
		summits = append(summits, summit)
	})
	if len(parsingErrors) != 0 {
		cumErrorString := ""
		for _, err := range parsingErrors {
			cumErrorString = fmt.Sprintf("%s, %s", cumErrorString, err.Error())
		}
		return nil, errors.New(cumErrorString)
	}
	return summits, nil
}
