package handler

import (
	"bufio"
	"fmt"
	"time"

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
