package handler

import (
	_ "effective_mobile_test_task/docs"
	"effective_mobile_test_task/internal/handler/api"
	"effective_mobile_test_task/internal/handler/api/base_api"
	"effective_mobile_test_task/internal/models"
	"effective_mobile_test_task/internal/service"
	"effective_mobile_test_task/pkg/misc"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var logger = misc.GetLogger()

func InitRoutes(service *service.Service, appConf *models.AppConfig, production bool) *gin.Engine {
	if production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	if !production {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	baseAPIRouter := baseApi.NewRouter(appConf)

	apiGroup := router.Group("/api/")
	apiRouter := api.NewRouter(baseAPIRouter, service)
	apiRouter.RegisterHandlers(apiGroup)

	return router
}
