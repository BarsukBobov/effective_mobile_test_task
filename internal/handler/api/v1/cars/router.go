package cars

import (
	baseApi "effective_mobile_test_task/internal/handler/api/base_api"
	"effective_mobile_test_task/internal/service"
	"github.com/gin-gonic/gin"
)

type router struct {
	*baseApi.Router
	service *service.CarsService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.CarsService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.GET("/get_all", h.getAll)
}
