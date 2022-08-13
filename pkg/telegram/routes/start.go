package routes

import (
	"context"
	"net/http"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type StartRequestBody struct{}

func Start(handler *telegram.Handler, logger *helper.Handler, client pb.TelegramServiceClient) {
	body := StartRequestBody{}
	ctx := handler.Context

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.Logger(err, &bindingCallback)

	res, err := client.Start(context.Background(), &pb.StartRequest{})
	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.Logger(err, &requestCallback)

	ctx.JSON(int(res.Status), &res)
}
