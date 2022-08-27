package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"go.uber.org/fx"
)

// EngineModule -> Dependency Injection for Gin Engine
var EngineModule = fx.Options(fx.Provide(NewHandler))

// Handler -> Dependency Injection Data Model for Gin Engine
type Handler struct {
	Gin *gin.Engine
}

/*
NewHandler -> Gin Engine initialization according to Mode
[config] -> Takes config as a parameter & checks {mode} for correct initialization

[return] -> returns Handler if config.Mode {correct} or {error}
*/
func NewHandler(config *config.Config) (h *Handler, err error) {
	if gin.DebugMode != config.Mode && gin.TestMode != config.Mode && gin.ReleaseMode != config.Mode {
		return nil, fmt.Errorf("incorrect mode type")
	}

	gin.SetMode(config.Mode)

	return &Handler{Gin: gin.Default()}, nil
}
