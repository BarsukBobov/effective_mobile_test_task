package v1

import (
	"effective_mobile_test_task/internal/handler/api/base_api"
	"effective_mobile_test_task/internal/handler/api/v1/cars"
	"effective_mobile_test_task/internal/service"
	"github.com/gin-gonic/gin"
)

type router struct {
	*baseApi.Router
	service *service.Service
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.Service,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	carsGroup := router.Group("/cars")
	carsRouter := cars.NewRouter(h.Router, h.service.Cars)
	carsRouter.RegisterHandlers(carsGroup)
}
