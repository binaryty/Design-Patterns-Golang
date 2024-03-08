package pkg

import (
	"fmt"
)

type YellowLight struct {
	tl *TrafficLight
}

func (l *YellowLight) GreenLight() error {
	return ErrNoWay
}

func (l *YellowLight) YellowLight() error {
	fmt.Println("Ready")

	Countdown(YellowDelay)

	fmt.Println()

	l.tl.setLight(l.tl.greenLight)

	return nil
}

func (l *YellowLight) RedLight() error {
	return ErrNoWay
}
