package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

// StopRequestBody -> Empty body
type StopRequestBody struct{}

/*
Stop -> Endpoint on telegram microservice

[ctx] ->
[logger] ->
[client] ->
*/
func Stop(ctx *gin.Context, logger *helper.LogHandler, client pb.TelegramServiceClient) {
	body := StopRequestBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.ErrorWithCallback(err, bindingCallback)

	res, err := client.Stop(context.Background(), &pb.StopRequest{})

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }
	logger.ErrorWithCallback(err, requestCallback)

	ctx.JSON(http.StatusOK, &res)
}
