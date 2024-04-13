package main

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Behavioral/State/traffic-light/pkg"
	"log"
)

func main() {
	tl := pkg.NewTrafficLight()

	for {
		if err := tl.GreenLight(); err != nil {
			log.Fatalf(err.Error())
		}

		if err := tl.YellowLight(); err != nil {
			log.Fatalf(err.Error())
		}

		if err := tl.RedLight(); err != nil {
			log.Fatalf(err.Error())
		}

		if err := tl.YellowLight(); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
