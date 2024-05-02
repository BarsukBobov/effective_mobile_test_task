package baseApi

type PaginationForm struct {
	Limit int `form:"limit" binding:"required"`
	Page  int `form:"page"  binding:"required"`
}
