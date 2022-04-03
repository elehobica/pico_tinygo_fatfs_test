package main

import (
	"fmt"
	"machine"
	"time"

	"mylocal.com/console"
	"tinygo.org/x/drivers/sdcard"
	"tinygo.org/x/tinyfs/fatfs"
)

var (
	spi    machine.SPI
	sckPin machine.Pin
	sdoPin machine.Pin
	sdiPin machine.Pin
	csPin  machine.Pin
	ledPin machine.Pin
)

func main() {
	sd := sdcard.New(spi, sckPin, sdoPin, sdiPin, csPin)
	err := sd.Configure()
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		for {
			time.Sleep(time.Hour)
		}
	}

	filesystem := fatfs.New(&sd)

	// Configure FATFS with sector size (must match value in ff.h - use 512)
	filesystem.Configure(&fatfs.Config{
		SectorSize: 512,
	})

	console.RunFor(&sd, filesystem)
}