package controller

import (
	"remote-box/controller/handler"

	"github.com/gin-gonic/gin"
)

func RouteController(port string) {
	dcc, err := handler.NewDCCEX("/dev/ttyACM0", 115200)
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	tracks := router.Group("/tracks")
	tracks.POST("/:trackName", dcc.SetTrack)

	vehicles := router.Group("/vehicles")
	vehicles.POST("/:dccAddress/speed", dcc.SetVehicleSpeed)
	vehicles.POST("/:dccAddress/func", dcc.SetVehicleFunction)

	ginErr := router.Run(port)
	if ginErr != nil {
		return
	}
}
