package telegram

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewHandler))

type Handler struct {
	Gin     *gin.Engine
	Context *gin.Context
}

func NewHandler() *Handler {
	handler := Handler{Gin: gin.Default(), Context: &gin.Context{}}
	return &handler
}
