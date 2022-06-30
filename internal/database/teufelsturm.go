package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jomaresch/climbDB/internal/database/models"
	"github.com/jomaresch/climbDB/internal/model"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TeufelsturmManager struct {
	db *sql.DB
}

func NewTeufelsturmManager(db *sql.DB) *TeufelsturmManager {
	return &TeufelsturmManager{db: db}
}

func (t *TeufelsturmManager) InsertRegion(ctx context.Context, region *model.Region, updateIfExists bool) error {
	ttRegion := &models.TTRegion{
		ID:          region.ID,
		DisplayName: region.DisplayName,
	}

	_, err := models.FindTTRegion(ctx, t.db, region.ID)
	if err == sql.ErrNoRows {
		zap.L().Debug("Region does not exists, insert new region.")
		return ttRegion.Insert(ctx, t.db, boil.Infer())
	}
	if err != nil {
		return fmt.Errorf("failed to check if region exists: %w", err)
	}

	if !updateIfExists {
		zap.L().Debug("Region already exists, skipping update.")
		return nil
	}

	zap.L().Debug("Region already exists, overwriting existing entry.")
	_, err = ttRegion.Update(ctx, t.db, boil.Infer())
	return err
}

func (t *TeufelsturmManager) ListRegions(ctx context.Context) ([]*model.Region, error) {
	regions, err := models.TTRegions().All(ctx, t.db)
	if err != nil {
		return nil, fmt.Errorf("failed to load regions: %w", err)
	}
	regionsSlice := make([]*model.Region, 0, len(regions))
	for _, reg := range regions {
		regionsSlice = append(regionsSlice, &model.Region{ID: reg.ID, DisplayName: reg.DisplayName})
	}
	return regionsSlice, nil
}

func (t *TeufelsturmManager) InsertSummit(ctx context.Context, summit *model.Summit, updateIfExists bool) error {
	ttSummit := &models.TTSummit{
		ID:          summit.ID,
		GuideID:     summit.GuideID,
		RegionID:    summit.RegionID,
		DisplayName: summit.DisplayName,
		Lat:         summit.Lat,
		Long:        summit.Long,
	}
	exists, err := models.TTSummitExists(ctx, t.db, summit.ID)
	if err != nil {
		return fmt.Errorf("failed to check if summit exists: %w", err)
	}
	if exists && !updateIfExists {
		zap.L().Debug("Summit already exists, skipping update.")
		return nil
	}
	if exists && updateIfExists {
		zap.L().Debug("Summit already exists, updating summit in db.")
		_, err = ttSummit.Update(ctx, t.db, boil.Infer())
	} else {
		zap.L().Debug("Insert new summit in db.")
		err = ttSummit.Insert(ctx, t.db, boil.Infer())
	}
	if err != nil {
		return fmt.Errorf("failed to update/insert summit: %w", err)
	}
	return nil
}

func (t *TeufelsturmManager) ListSummits(ctx context.Context) ([]*model.Summit, error) {
	summitsSlice, err := models.TTSummits().All(ctx, t.db)
	if err != nil {
		return nil, fmt.Errorf("failed list all summits: %w", err)
	}
	summits := make([]*model.Summit, 0, len(summitsSlice))
	for _, summit := range summitsSlice {
		summits = append(summits, &model.Summit{
			ID:          summit.ID,
			GuideID:     summit.GuideID,
			RegionID:    summit.RegionID,
			DisplayName: summit.DisplayName,
			Lat:         summit.Lat,
			Long:        summit.Long,
		})
	}
	return summits, nil
}

func (t *TeufelsturmManager) InsertRoute(ctx context.Context, route *model.Route, updateIfExists bool) error {
	l := zap.L().With(zap.String("id", route.ID), zap.String("name", route.DisplayName))

	ttRoute := &models.TTRoute{
		ID:               route.ID,
		SummitID:         route.SummitID,
		DisplayName:      route.DisplayName,
		AvgRating:        int64(route.AvgRating),
		GuideRating:      int64(route.GuideRating),
		GradeID:          int64(route.GradeID),
		SuggestedGradeID: int64(route.SuggestedGradeID),
		RedPointGradeID:  int64(route.RedPointGradeID),
		JumpGrade:        int64(route.JumpGrade),
	}

	exists, err := models.TTRouteExists(ctx, t.db, route.ID)
	if err != nil {
		return fmt.Errorf("failed to check if route exists: %w", err)
	}
	if exists && !updateIfExists {
		l.Debug("Route already exists, skipping update.")
		return nil
	}
	if exists && updateIfExists {
		l.Debug("Route already exists, updating route in db.")
		_, err = ttRoute.Update(ctx, t.db, boil.Infer())
	} else {
		l.Debug("Insert new route in db.")
		err = ttRoute.Insert(ctx, t.db, boil.Infer())
	}
	if err != nil {
		return fmt.Errorf("failed to update/insert route: %w", err)
	}
	return nil
}

func (t *TeufelsturmManager) ListRoutes(ctx context.Context) (map[string]*model.Route, error) {
	routesSlice, err := models.TTRoutes().All(ctx, t.db)
	if err != nil {
		return nil, fmt.Errorf("failed list all routes: %w", err)
	}
	routes := make(map[string]*model.Route, len(routesSlice))
	for _, route := range routesSlice {
		routes[route.ID] = &model.Route{
			ID:               route.ID,
			SummitID:         route.SummitID,
			DisplayName:      route.DisplayName,
			AvgRating:        int(route.AvgRating),
			GuideRating:      int(route.GuideRating),
			GradeID:          int(route.GradeID),
			SuggestedGradeID: int(route.SuggestedGradeID),
			RedPointGradeID:  int(route.RedPointGradeID),
			JumpGrade:        int(route.JumpGrade),
		}
	}
	return routes, nil
}
