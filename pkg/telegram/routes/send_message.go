package routes

import (
	"context"
	"net/http"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type SendMessageBody struct {
	Message string `json:"message"`
}

func SendMessage(handler *telegram.Handler, logger *helper.Handler, client pb.TelegramServiceClient) {
	body := SendMessageBody{}
	ctx := handler.Context

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.Logger(err, &bindingCallback)

	res, err := client.SendMessage(
		context.Background(), &pb.SendMessageRequest{Message: body.Message},
	)

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.Logger(err, &requestCallback)

	ctx.JSON(int(res.Status), &res)
}
