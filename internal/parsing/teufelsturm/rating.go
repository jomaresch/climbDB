package teufelsturm

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseRatingFromPictureElement(element *goquery.Selection) (int, error) {
	sourceReference, found := element.Attr("src")
	if !found {
		return 0, fmt.Errorf("no 'src' attribute found")
	}

	switch {
	case strings.HasSuffix(sourceReference, "arrow-downright3.gif"):
		return -3, nil
	case strings.HasSuffix(sourceReference, "arrow-downright2.gif"):
		return -2, nil
	case strings.HasSuffix(sourceReference, "arrow-downright.gif"):
		return -1, nil
	case strings.HasSuffix(sourceReference, "arrow-right.gif"):
		return 0, nil
	case strings.HasSuffix(sourceReference, "arrow-upright.gif"):
		return 1, nil
	case strings.HasSuffix(sourceReference, "arrow-upright2.gif"):
		return 2, nil
	case strings.HasSuffix(sourceReference, "arrow-upright3.gif"):
		return 3, nil
	}
	return 0, fmt.Errorf("unable to parse rating icon: %s", sourceReference)
}

func ParseRatingFromText(ratingText string) int {
	switch {
	case strings.HasPrefix(ratingText, "--- "):
		return -3
	case strings.HasPrefix(ratingText, "-- "):
		return -2
	case strings.HasPrefix(ratingText, "- "):
		return -1
	case strings.HasPrefix(ratingText, "+++ "):
		return 3
	case strings.HasPrefix(ratingText, "++ "):
		return 2
	case strings.HasPrefix(ratingText, "+ "):
		return 1
	default:
		return 0
	}
}
