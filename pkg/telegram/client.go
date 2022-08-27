package telegram

import (
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientModule -> Dependency Injection for Client module
var ClientModule = fx.Options(fx.Provide(InitServiceClient), fx.Provide(Initialize))

// ServiceClient -> Dependency Injection Data Model for ClientModule
type ServiceClient struct {
	Client pb.TelegramServiceClient
	Logger *helper.LogHandler
}

/*
Initialize -> ServiceClient initialization that should be used for Dependency Injection

[client] -> {TelegramServiceClient} using generated with {grpc client}
[logger] -> General LogHandler to log errors

[returns] -> {ServiceClient} generated with {TelegramServiceClient} using generated with {grpc client} & {logger}
*/
func Initialize(client pb.TelegramServiceClient, logger *helper.LogHandler) ServiceClient {
	return ServiceClient{
		Client: client,
		Logger: logger,
	}
}

/*
InitServiceClient -> Generate grpc client for {Api Gateway}

[return] -> returns {TelegramServiceClient} using generated with {grpc client}
*/
func InitServiceClient(c *config.Config, logger *helper.LogHandler) pb.TelegramServiceClient {
	cc, err := grpc.Dial(c.TelegramSvcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	logger.Error(err)

	return pb.NewTelegramServiceClient(cc)
}
