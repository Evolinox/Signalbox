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

	router.POST("/tracks/main/on", func(c *gin.Context) {
		err := dcc.Send("1")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "on"})
	})

	router.POST("/tracks/main/off", func(c *gin.Context) {
		err := dcc.Send("0")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "off"})
	})

	ginErr := router.Run(port)
	if ginErr != nil {
		return
	}
}
