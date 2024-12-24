package main

import (
	"log"

	"source-crdc.hirain.com/atx-go/go-socketcan/pkg/socketcan"
)

func main() {
	//发送canfd报文
	output, err := socketcan.NewRawInterface("can1")
	if err != nil {
		log.Fatalf("Failed to open %s: %v", "can1", err)
	}
	msg := socketcan.CanFrame{
		ArbId:    0x12345678,
		Dlc:      8,
		Data:     []byte{0x04, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00},
		Extended: true,
		//根据IsFD发送对应报文
		IsFD: true,
	}
	if err := output.SendFrame(msg); err != nil {
		log.Printf("Error transmitting frame : %v", err)
	}

}
