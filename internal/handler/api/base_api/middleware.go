package baseApi

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (h *Middleware) SessionRequired(c *gin.Context) {
}

func (h *Middleware) AdminRequired(c *gin.Context) {
}

func (h *Middleware) AuthRequired(c *gin.Context) {
}
