package api

import (
	"effective_mobile_test_task/internal/handler/api/base_api"
	"effective_mobile_test_task/internal/handler/api/v1"
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
	v1Group := router.Group("/v1")
	v1Router := v1.NewRouter(h.Router, h.service)
	v1Router.RegisterHandlers(v1Group)
}
