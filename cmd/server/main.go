package main

import (
	"context"
	"database/sql"

	"github.com/go-resty/resty/v2"
	"github.com/jomaresch/climbDB/internal/crawling"
	"github.com/jomaresch/climbDB/internal/database"
	"github.com/jomaresch/climbDB/internal/model"
	"go.uber.org/zap"
)

func init() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	l, _ := cfg.Build()
	zap.ReplaceGlobals(l)
}

func main() {
	db, err := sql.Open("sqlite3", "./assets/climb.db")
	fatalOnError("failed to open database", err)
	defer db.Close()
	m := database.NewTeufelsturmManager(db)

	client := resty.New()
	teufelsturmCrawler := crawling.NewTeufelsturmCrawler(client)

	storeAllComments(teufelsturmCrawler, m)

	//regionsMapping := model.MapRegionsNames(regions)
	//summits, err := teufelsturmCrawler.ExtractAllSummits(context.Background(), regionsMapping)
	//fatalOnError("failed to crawl summits", err)
	//summitsMapping := model.MapSummitNames(summits)
	//routes, err := teufelsturmCrawler.ExtractAllRoutes(context.Background(), summitsMapping)
	//fatalOnError("failed to crawl routes", err)
	//comments := make([]*model.Comment, 0, len(routes))
	//for _, route := range routes {
	//	c, err := teufelsturmCrawler.ExtractComments(context.Background(), route.ID)
	//	fatalOnError("failed to load route details", err)
	//	comments = append(comments, c...)
	//	fmt.Printf("%s === %d \n", route.ID, len(comments))
	//}
	//fmt.Println(comments)
}

func storeAllRegions(tc *crawling.TeufelsturmCrawler, tm *database.TeufelsturmManager) {
	regions, err := tc.ExtractAllRegions(context.Background())
	fatalOnError("failed to crawl regions", err)
	for _, region := range regions {
		err = tm.InsertRegion(context.Background(), region, true)
		fatalOnError("failed to store region", err)
	}
}

func loadAllRegions(tm *database.TeufelsturmManager) map[string]string {
	regions, err := tm.ListRegions(context.Background())
	fatalOnError("failed to load regions", err)
	return model.MapRegionsNames(regions)
}

func storeAllSummits(tc *crawling.TeufelsturmCrawler, tm *database.TeufelsturmManager) {
	summits, err := tc.ExtractAllSummits(context.Background(), loadAllRegions(tm), true)
	fatalOnError("failed to crawl summits", err)
	for _, summit := range summits {
		err = tm.InsertSummit(context.Background(), summit, true)
		fatalOnError("failed to store summit", err)
	}
}

func storeAllRoutes(tc *crawling.TeufelsturmCrawler, tm *database.TeufelsturmManager) {
	routes, err := tc.ExtractAllRoutes(context.Background(), loadAllSummits(tm))
	fatalOnError("failed to crawl routes", err)
	for _, route := range routes {
		err = tm.InsertRoute(context.Background(), route, true)
		fatalOnError("failed to store route", err)
	}
}

func storeAllComments(tc *crawling.TeufelsturmCrawler, tm *database.TeufelsturmManager) {
	routesMap, err := tm.ListRoutes(context.Background())
	fatalOnError("failed to load all routes", err)
	count := 0
	for key := range routesMap {
		count++
		comments, err := tc.ExtractComments(context.Background(), key)
		if count%100 == 0 {
			zap.S().Info(float64(count) / float64(len(routesMap)))
		}
		fatalOnError("failed to crawl comments for route", err)
		for _, comment := range comments {
			err = tm.InsertComment(context.Background(), comment, true)
			fatalOnError("failed to store comment", err)
		}
	}
}

func loadAllSummits(tm *database.TeufelsturmManager) map[string]string {
	summits, err := tm.ListSummits(context.Background())
	fatalOnError("failed to load summits", err)
	return model.MapSummitNames(summits)
}

func fatalOnError(msg string, err error) {
	if err != nil {
		zap.L().Fatal(msg, zap.Error(err))
	}
}
