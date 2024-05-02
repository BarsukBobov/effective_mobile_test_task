package cars

import (
	baseApi "effective_mobile_test_task/internal/handler/api/base_api"
	"effective_mobile_test_task/internal/repository/sql"
)

type createForm struct {
	RegNums []string `json:"regNums" binding:"required"`
}

type editForm struct {
	*sql.EditCar
}

type getAllForm struct {
	baseApi.PaginationForm
	*sql.CarFilter
}

type getAllResponse struct {
	Data  []sql.Car `json:"data"`
	Total int       `json:"total"`
}
