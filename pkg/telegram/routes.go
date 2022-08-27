package telegram

import (
	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/routes"
	"go.uber.org/fx"
)

// RouteModule -> Dependency Injection for RouteModule module
var RouteModule = fx.Options(fx.Invoke(RegisterRoutes))

/*
RegisterRoutes -> Register Telegram service routes to Gin

[client] -> Takes arguments as Service client
[handler] -> Takes arguments as Gin.Engine
*/
func RegisterRoutes(client ServiceClient, handler *pkg.Handler) {
	routes := handler.Gin.Group("/telegram")
	routes.POST("/", client.Start)
	routes.POST("/healthz", client.Status)
	routes.POST("/message", client.SendMessage)
	routes.POST("/stop", client.Stop)

	// TODO: will be deleted preserved for test purposes only
	routes.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Health OK"})
	})
}

// Start -> Telegrams Service endpoint
func (svc ServiceClient) Start(ctx *gin.Context) {
	routes.Start(ctx, svc.Logger, svc.Client)
}

// Stop -> Telegrams Service endpoint
func (svc ServiceClient) Stop(ctx *gin.Context) {
	routes.Stop(ctx, svc.Logger, svc.Client)
}

// SendMessage -> Telegrams Service endpoint
func (svc ServiceClient) SendMessage(ctx *gin.Context) {
	routes.SendMessage(ctx, svc.Logger, svc.Client)
}

// Status -> Telegrams Service endpoint
func (svc ServiceClient) Status(ctx *gin.Context) {
	routes.Status(ctx, svc.Logger, svc.Client)
}
