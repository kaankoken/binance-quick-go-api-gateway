package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type StatusMessageBody struct{}

func Status(ctx *gin.Context, logger *helper.LogHandler, client pb.TelegramServiceClient) {
	body := StatusMessageBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.ErrorWithCallback(err, bindingCallback)

	res, err := client.Status(context.Background(), &pb.StatusRequest{})

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.ErrorWithCallback(err, requestCallback)

	ctx.JSON(http.StatusOK, &res)
}
