package main

import (
	"log"

	"wm-motor.com/Infra/cobratest/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Execute err: %v", err)
	}
}
