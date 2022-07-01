package crawling

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jomaresch/climbDB/internal/model"
	"github.com/jomaresch/climbDB/internal/parsing/teufelsturm"
)

const (
	regionsListUrl         = "http://teufelsturm.de/gebiete"
	summitsSearchUrl       = "http://teufelsturm.de/gipfel/suche.php"
	summitDetailsUrlFormat = "http://teufelsturm.de/gipfel/details.php?gipfelnr=%s"
	routesSearchUrl        = "http://teufelsturm.de/wege/suche.php"
	routeDetailsUrlFormat  = "http://teufelsturm.de/wege/bewertungen/anzeige.php?wegnr=%s"
)

type TeufelsturmCrawler struct {
	client *resty.Client
}

func NewTeufelsturmCrawler(client *resty.Client) *TeufelsturmCrawler {
	return &TeufelsturmCrawler{client: client}
}

func (t *TeufelsturmCrawler) ExtractAllRegions(ctx context.Context) ([]*model.Region, error) {
	resp, err := t.client.NewRequest().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		Get(regionsListUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get regions list page: %w", err)
	}
	regions, err := teufelsturm.ParseRegionList(resp.RawBody())
	if err != nil {
		return nil, fmt.Errorf("failed to parse regions list page: %w", err)
	}
	return regions, nil
}

func (t *TeufelsturmCrawler) ExtractAllSummits(ctx context.Context, regionsMapping map[string]string, withGPS bool) ([]*model.Summit, error) {
	return t.extractSummitsByQuery(ctx, map[string]string{"gebietnr": "0", "sortiert": "2", "anzahl": "3000"}, regionsMapping, withGPS)
}

func (t *TeufelsturmCrawler) extractSummitsByQuery(ctx context.Context, query, regionsMapping map[string]string, withGPS bool) ([]*model.Summit, error) {
	resp, err := t.client.NewRequest().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		SetFormData(query).
		Post(summitsSearchUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to search for summits: %w", err)
	}
	summits, err := teufelsturm.ParseSummitList(resp.RawBody(), regionsMapping)
	if err != nil {
		return nil, fmt.Errorf("failed to parse summits list page: %w", err)
	}
	if !withGPS {
		return summits, nil
	}

	for _, summit := range summits {
		resp, err = t.client.NewRequest().
			SetContext(ctx).
			SetDoNotParseResponse(true).
			Get(fmt.Sprintf(summitDetailsUrlFormat, summit.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to load details for summit '%s': %w", summit.ID, err)
		}

		summitWithGPS, err := teufelsturm.ParseSummitDetails(resp.RawBody())
		if err != nil {
			return nil, fmt.Errorf("failed to parse details for summit '%s': %w", summit.ID, err)
		}
		summit.Lat = summitWithGPS.Lat
		summit.Long = summitWithGPS.Long
	}
	return summits, nil
}

func (t *TeufelsturmCrawler) ExtractAllRoutes(ctx context.Context, summitsMapping map[string]string) ([]*model.Route, error) {
	return t.extractRoutesByQuery(ctx, map[string]string{
		"text":         "",
		"gebiet":       "0",
		"skala_von":    "",
		"skala_bis":    "",
		"benutzer":     "",
		"bewertungen":  "",
		"avgbewertung": "",
		"sortiert":     "1",
		"datum":        "1",
		"anzahl":       "20000",
	}, summitsMapping)
}

func (t *TeufelsturmCrawler) extractRoutesByQuery(ctx context.Context, query, summitsMapping map[string]string) ([]*model.Route, error) {
	resp, err := t.client.NewRequest().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		SetFormData(query).
		Post(routesSearchUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to search for routes: %w", err)
	}
	routes, err := teufelsturm.ParseRouteList(resp.RawBody(), summitsMapping)
	if err != nil {
		return nil, fmt.Errorf("failed to parse routes list page: %w", err)
	}
	return routes, nil
}

func (t *TeufelsturmCrawler) ExtractComments(ctx context.Context, routeId string) ([]*model.Comment, error) {
	resp, err := t.client.NewRequest().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf(routeDetailsUrlFormat, routeId))
	if err != nil {
		return nil, fmt.Errorf("failed to get route details: %w", err)
	}
	comments, err := teufelsturm.ParseRouteDetails(resp.RawBody(), routeId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse route details: %w", err)
	}
	return comments, nil
}
