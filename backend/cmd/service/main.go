package main

import (
	"sn-go-api/internal/config"
	"sn-go-api/internal/router"
)

func main() {
	snConfig := config.Init()
	r := router.SetupRouter(snConfig)
	r.Run(":8080")
}
