package pkg

import (
	"fmt"
)

type RedLight struct {
	tl *TrafficLight
}

func (r RedLight) GreenLight() error {
	return ErrNoWay
}

func (r RedLight) YellowLight() error {
	return ErrNoWay
}

func (r RedLight) RedLight() error {
	fmt.Println("Stop")

	Countdown(RedDelay)

	fmt.Println()

	r.tl.setLight(r.tl.yellowLight)

	return nil
}
