package main

import (
	"fmt"

	"order_calculator/app/api"
	"order_calculator/app/config"
)

func main() {
	cfg := config.LoadConfig()
	r := api.NewRouter(cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Port()))
}
