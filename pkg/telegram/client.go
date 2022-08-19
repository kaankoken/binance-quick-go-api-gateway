package telegram

import (
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ClientModule = fx.Options(fx.Provide(InitServiceClient))

type ServiceClient struct {
	Client pb.TelegramServiceClient
	Logger *helper.LogHandler
}

func InitServiceClient(c *config.Config, logger *helper.LogHandler) pb.TelegramServiceClient {
	cc, err := grpc.Dial(c.TelegramSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	logger.Error(err)

	return pb.NewTelegramServiceClient(cc)
}
