package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

// StartRequestBody -> Empty body
type StartRequestBody struct{}

/*
Start -> Endpoint on telegram microservice

[ctx] ->
[logger] ->
[client] ->
*/
func Start(ctx *gin.Context, logger *helper.LogHandler, client pb.TelegramServiceClient) {
	body := StartRequestBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.ErrorWithCallback(err, bindingCallback)

	res, err := client.Start(context.Background(), &pb.StartRequest{})
	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.ErrorWithCallback(err, requestCallback)

	ctx.JSON(int(res.Status), &res)
}
