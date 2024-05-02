package baseApi

import (
	"effective_mobile_test_task/internal/models"
	"effective_mobile_test_task/pkg/misc"
	"github.com/gin-gonic/gin"
)

var logger = misc.GetLogger()

type Router struct {
	Middleware *Middleware
	AppConf    *models.AppConfig
}

func NewRouter(appConf *models.AppConfig) *Router {
	return &Router{
		Middleware: NewMiddleware(),
		AppConf:    appConf,
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
