package api

import (
	"net/http"
	"order_calculator/app/calculator"

	"github.com/gin-gonic/gin"
)

type GetPacksResponse struct {
	PackSizes []int `json:"pack_sizes"`
}

func getPackSizes(availablePackSizes []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, GetPacksResponse{PackSizes: availablePackSizes})
	}
}

type PostOrderResponse struct {
	OrderConfiguration map[int]int `json:"order_configuration,omitempty"`
	Error              string      `json:"error,omitempty"`
}

func calculateOrder(availablePackSizes []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response PostOrderResponse
		var req OrderRequest
		if err := ctx.ShouldBind(&req); err != nil {
			response.Error = err.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		if req.OrderSize < 1 {
			response.Error = "invalid order size, must be bigger than 0"
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		response.OrderConfiguration = calculator.CalculateOrderConfiguration(availablePackSizes, req.OrderSize)

		ctx.JSON(http.StatusOK, response)
	}
}
