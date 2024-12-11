package main

import (
	"github.com/Philipsco/pbk-mrt-schedules/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	initRouter()
}

func initRouter() {

	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)
	router.Run(":8080")
}
