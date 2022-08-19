package telegram

import (
	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/routes"
	"go.uber.org/fx"
)

var RouteModule = fx.Options(fx.Provide(initialize), fx.Invoke(registerRoutes))

func initialize(client pb.TelegramServiceClient, logger *helper.LogHandler) *ServiceClient {
	svc := &ServiceClient{
		Client: client,
		Logger: logger,
	}

	return svc
}

func registerRoutes(client *ServiceClient, handler *pkg.Handler) {
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
