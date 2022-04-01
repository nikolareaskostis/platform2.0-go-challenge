package main

import (
	"GoRepo/source/assignment/cache"
	controllers "GoRepo/source/assignment/controllers"
	"GoRepo/source/assignment/dataAccess"
	"GoRepo/source/assignment/db"
	"GoRepo/source/assignment/environment"
	"GoRepo/source/assignment/middlewares"
	"GoRepo/source/assignment/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func main() {
	Run()
}

func Run() {

	//Load Configuration Variables
	config := environment.LoadConfig("config.yml")

	cache := cache.CacheInit()

	//Initialize db - Connect and Build Schema
	dbConn, err := setUpDatabase(config)

	if err != nil {
		//handel failure in db initialization
	}

	//Set up Controllers & endpoints - Start serving
	router := setUpServer(config, dbConn, cache)
	router.Run(config.Server.Port)
}

func setUpDatabase(config environment.Config) (gorm.DB, error) {
	dbConn, err := db.GetGormConnection(config.Database.ConnectionString)

	return *dbConn, err
}

func setUpServer(config environment.Config, dbConn gorm.DB, cache *redis.Client) *gin.Engine {

	//Initialize DataAccess - Services -Controllers
	authorizeDa := dataAccess.NewAuthorize_da(dbConn)
	authorizeService := services.NewAuthService(authorizeDa)
	tokenService := services.NewTokenService(cache)
	authorizeController := controllers.NewAuthorizeController(authorizeService, tokenService)

	chartDa := dataAccess.NewChart_da(dbConn)
	chartService := services.NewChartService(chartDa)
	chartController := controllers.NewChartsController(chartService)

	assetDa := dataAccess.NewAsset_da(dbConn)
	assetService := services.NewAssetService(assetDa)
	assetController := controllers.NewAssetsController(assetService)

	audienceDa := dataAccess.NewAudience_da(dbConn)
	audienceService := services.NewAudienceService(audienceDa)
	audienceController := controllers.NewAudienceController(audienceService)

	insightDa := dataAccess.NewInsight_da(dbConn)
	insightService := services.NewInsightService(insightDa)
	insightController := controllers.NewInsightController(insightService)

	// Initialize Middlewares
	middleware := middlewares.NewMiddleware(tokenService)

	router := gin.Default()

	// Setup CORS //todo na perasw ta ports apo to env config
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:7788", "http://localhost:2200"},
		AllowMethods:  []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// No authorization endpoints
	controller := router.Group("/api/controllers")
	{
		controller.POST("authorize/login", authorizeController.Login)
		controller.POST("authorize/register", authorizeController.Register)
		controller.GET("authorize/check", authorizeController.Check)
	}

	// Authorized endpoints - using middleware with http handlers and jwt tokens to authorize
	controller.Use(middleware.ValidateJwtToken())
	{
		controller.GET("asset/get", assetController.Get)
		controller.GET("asset/getfavourites", assetController.GetFavourites)

		controller.GET("chart/get", chartController.Get)
		controller.GET("chart/getfavourites", chartController.GetFavourites)
		controller.GET("chart/getchart", chartController.GetChart)
		controller.PUT("chart/setfavourite", chartController.SetFavourite)
		controller.PUT("chart/unsetfavourite", chartController.UnsetFavourite)
		controller.DELETE("chart/delete", chartController.Delete)

		controller.GET("audience/get", audienceController.Get)
		controller.GET("audience/getfavourites", audienceController.GetFavourites)
		controller.PUT("audience/setfavourite", audienceController.SetFavourite)
		controller.PUT("audience/unsetfavourite", audienceController.UnsetFavourite)
		controller.DELETE("audience/delete", audienceController.Delete)

		controller.GET("insight/get", insightController.Get)
		controller.GET("insight/getfavourites", insightController.GetFavourites)
		controller.PUT("insight/setfavourite", insightController.SetFavourite)
		controller.PUT("insight/unsetfavourite", insightController.UnsetFavourite)
		controller.DELETE("insight/delete", insightController.Delete)
	}

	return router
}
