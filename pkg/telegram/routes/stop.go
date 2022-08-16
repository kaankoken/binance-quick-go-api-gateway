package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
)

type StopRequestBody struct{}

func Stop(ctx *gin.Context, logger *helper.Handler, client pb.TelegramServiceClient) {
	body := StopRequestBody{}

	err := ctx.BindJSON(&body)

	bindingCallback := func() { ctx.AbortWithError(http.StatusBadRequest, err) }
	logger.Error(err, &bindingCallback)

	res, err := client.Stop(context.Background(), &pb.StopRequest{})

	requestCallback := func() { ctx.AbortWithError(http.StatusBadGateway, err) }
	logger.Error(err, &requestCallback)

	ctx.JSON(http.StatusOK, &res)
}
