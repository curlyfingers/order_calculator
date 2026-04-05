package api

import (
	"order_calculator/app/config"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	OrderSize int `form:"order_size" json:"order_size"`
}

func NewRouter(cfg config.Config) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/packs", getPackSizes(cfg.PackSizes()))
	api.POST("/order", calculateOrder(cfg.PackSizes()))

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("./client/index.html")
	})

	router.GET("/assets/script.js", func(ctx *gin.Context) {
		ctx.File("./client/script.js")
	})

	return router
}
