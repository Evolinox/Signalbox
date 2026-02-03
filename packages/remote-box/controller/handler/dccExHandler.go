package handler

import (
	"bufio"
	"fmt"
	"remote-box/entities"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tarm/serial"
)

type DCCEX struct {
	port *serial.Port
}

func NewDCCEX(portName string, baud int) (*DCCEX, error) {
	c := &serial.Config{Name: portName, Baud: baud, ReadTimeout: time.Second * 2}
	p, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	return &DCCEX{port: p}, nil
}

func (d *DCCEX) Send(cmd string) error {
	full := fmt.Sprintf("<%s>\n", cmd)
	_, err := d.port.Write([]byte(full))
	return err
}

func (d *DCCEX) Read() (string, error) {
	reader := bufio.NewReader(d.port)
	line, err := reader.ReadString('>')
	return line, err
}

/*
boolToInt - Small helper function to convert bools to either 0 or 1
*/
func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

/*
SetTrack - Function to set the powerstate of MAIN/PROG Tracks
*/
func (d *DCCEX) SetTrack(c *gin.Context) {
	trackName := c.Param("trackName")
	var req entities.TrackRequest

	// Bind Requestbody and Check if invalid
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Format DCC-EX Command
	dccCmd := fmt.Sprintf("<%d %s>", boolToInt(req.Power), strings.ToUpper(trackName))

	if err := d.Send(dccCmd); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

/*
SetVehicleSpeed
*/
func (d *DCCEX) SetVehicleSpeed(c *gin.Context) {
	dccAddress := c.Param("dccAddress")
	var req entities.VehicleSpeedRequest

	// Bind Requestbody and Check if invalid
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Format DCC-EX Command
	dccCmd := fmt.Sprintf("<F %s %d %d>", dccAddress, req.Speed, boolToInt(req.Forward))

	if err := d.Send(dccCmd); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

/*
SetVehicleFunction - Needs an DCC Address, Function Number and its State (0 = Off, 1 = On)
*/
func (d *DCCEX) SetVehicleFunction(c *gin.Context) {
	dccAddress := c.Param("dccAddress")
	var req entities.VehicleFunctionRequest

	// Bind Requestbody and Check if invalid
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Format DCC-EX Command
	dccCmd := fmt.Sprintf("<F %s %d %d>", dccAddress, req.Function, req.State)

	if err := d.Send(dccCmd); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
