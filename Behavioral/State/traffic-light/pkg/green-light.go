package pkg

import (
	"fmt"
)

type GreenLight struct {
	tl *TrafficLight
}

func (g GreenLight) GreenLight() error {
	fmt.Println("Go")

	Countdown(GreenDelay)

	fmt.Println()
	g.tl.setLight(g.tl.yellowLight)
	g.tl.prevLight = g.tl.greenLight

	return nil
}

func (g GreenLight) YellowLight() error {
	return ErrNoWay
}

func (g GreenLight) RedLight() error {
	return ErrNoWay
}
