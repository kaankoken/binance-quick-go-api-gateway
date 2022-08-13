package routes

import (
	"context"
	"net/http"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type StatusMessageBody struct{}

func Status(handler *telegram.Handler, logger *helper.Handler, client pb.TelegramServiceClient) {
	body := StatusMessageBody{}
	ctx := handler.Context

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.Logger(err, &bindingCallback)

	res, err := client.Status(context.Background(), &pb.StatusRequest{})

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.Logger(err, &requestCallback)

	ctx.JSON(http.StatusOK, &res)
}
