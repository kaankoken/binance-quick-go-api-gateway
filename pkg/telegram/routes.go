package telegram

import (
	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/routes"
	"go.uber.org/fx"
)

var RouteModule = fx.Options(fx.Invoke(RegisterRoutes))

func RegisterRoutes(handler *Handler, config *config.Config, logger *helper.Handler) {
	svc := &ServiceClient{
		Client: InitServiceClient(config, logger),
		Logger: logger,
	}

	routes := handler.Gin.Group("/telegram")
	routes.POST("/", svc.Start)
	routes.POST("/healthz", svc.Status)
	routes.POST("/message", svc.SendMessage)
	routes.POST("/stop", svc.Stop)

	// TODO: will be deleted preserved for test purposes only
	routes.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Health OK"})
	})
}

func (svc *ServiceClient) Start(ctx *gin.Context) {
	routes.Start(ctx, svc.Logger, svc.Client)
}

func (svc *ServiceClient) Stop(ctx *gin.Context) {
	routes.Stop(ctx, svc.Logger, svc.Client)
}

func (svc *ServiceClient) SendMessage(ctx *gin.Context) {
	routes.SendMessage(ctx, svc.Logger, svc.Client)
}

func (svc *ServiceClient) Status(ctx *gin.Context) {
	routes.Status(ctx, svc.Logger, svc.Client)
}
