package cars

import (
	"effective_mobile_test_task/internal/handler/api/base_api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary create
// @Tags cars
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object} baseApi.SuccessResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /cars/ [post]
func (h *router) create(c *gin.Context) {
	var form createForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	err = h.service.Create(form.RegNums)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.NewSuccessResponse(c)
}

// @Summary edit
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object} baseApi.SuccessResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /cars/{id} [put]
func (h *router) edit(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	var form editForm
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	err = h.service.Edit(id, form.EditCar)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.NewSuccessResponse(c)
}

// @Summary delete
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} baseApi.SuccessResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /cars/{id} [delete]
func (h *router) delete(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.NewSuccessResponse(c)
}

// @Summary getAll
// @Tags cars
// @Accept json
// @Produce json
// @Param q query getAllForm true "filter car form"
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /cars/get_all [get]
func (h *router) getAll(c *gin.Context) {
	var form getAllForm
	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	total, cars, err := h.service.GetAll(form.Limit, form.Page, form.CarFilter)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, getAllResponse{Data: cars, Total: total})
}
