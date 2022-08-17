package telegram

import (
	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"go.uber.org/fx"
)

var EngineModule = fx.Options(fx.Provide(NewHandler))

type Handler struct {
	Gin *gin.Engine
}

func NewHandler(config *config.Config) *Handler {
	gin.SetMode(config.Mode)

	handler := Handler{Gin: gin.Default()}
	return &handler
}
