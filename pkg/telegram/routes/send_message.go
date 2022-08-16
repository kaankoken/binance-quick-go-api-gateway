package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type SendMessageBody struct {
	Message string `json:"message"`
}

func SendMessage(ctx *gin.Context, logger *helper.Handler, client pb.TelegramServiceClient) {
	body := SendMessageBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.Error(err, &bindingCallback)

	res, err := client.SendMessage(
		context.Background(), &pb.SendMessageRequest{Message: body.Message},
	)

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.Error(err, &requestCallback)

	ctx.JSON(int(res.Status), &res)
}
