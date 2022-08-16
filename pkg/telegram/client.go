package telegram

import (
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.TelegramServiceClient
	Logger *helper.Handler
}

func InitServiceClient(c *config.Config, logger *helper.Handler) pb.TelegramServiceClient {
	cc, err := grpc.Dial(c.TelegramSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	logger.Error(err, nil)

	return pb.NewTelegramServiceClient(cc)
}
