package main

import "source-crdc.hirain.com/atx-go/go-socketcan/pkg/candevice"

func main() {
	// Error handling omitted to keep example simple
	d, _ := candevice.New("can0")
	_ := d.SetBitrate(250000)
	_ := d.SetUp()
	defer d.SetDown()
}