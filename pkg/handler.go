package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"go.uber.org/fx"
)

var EngineModule = fx.Options(fx.Provide(NewHandler))

type Handler struct {
	Gin *gin.Engine
}

func NewHandler(config *config.Config) (h *Handler, err error) {
	if gin.DebugMode != config.Mode && gin.TestMode != config.Mode && gin.ReleaseMode != config.Mode {
		return nil, fmt.Errorf("incorrect mode type")
	}

	gin.SetMode(config.Mode)

	return &Handler{Gin: gin.Default()}, nil
}
