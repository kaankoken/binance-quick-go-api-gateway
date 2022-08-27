package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

/*
SendMessageBody -> body of the SendMessage endpoint
[Message] -> Takes string in the body
*/
type SendMessageBody struct {
	Message string `json:"message"`
}

/*
SendMessage -> Endpoint on telegram microservice

[ctx] ->
[logger] ->
[client] ->
*/
func SendMessage(ctx *gin.Context, logger *helper.LogHandler, client pb.TelegramServiceClient) {
	body := SendMessageBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.ErrorWithCallback(err, bindingCallback)

	res, err := client.SendMessage(
		context.Background(), &pb.SendMessageRequest{Message: body.Message},
	)

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }

	logger.ErrorWithCallback(err, requestCallback)

	ctx.JSON(int(res.Status), &res)
}
