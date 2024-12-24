package main

import (
	"log"

	"source-crdc.hirain.com/atx-go/go-socketcan/pkg/socketcan"
)

func main() {

	//不处理err，如果要处理err后面加
	// if err != nil {
	// 	log.Fatalf("Failed to open %s: %v", "can1", err)
	// }

	input, _ := socketcan.NewRawInterface("can1")
	output, _ := socketcan.NewRawInterface("can2")

	for {
		//收报文
		frame, err := input.RecvFrame()
		if err != nil {
			log.Printf("Error receiving frame: %v", err)
			continue
		}
		//发报文
		output.SendFrame(frame)
	}
}
