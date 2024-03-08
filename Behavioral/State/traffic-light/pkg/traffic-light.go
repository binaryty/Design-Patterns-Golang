package pkg

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	GreenDelay  = "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15"
	RedDelay    = "1 2 3 4 5 6 7 8 9 10"
	YellowDelay = "1 2 3"
)

var (
	ErrNoWay = errors.New("no way")
)

type TrafficLight struct {
	currentLight Light
	greenLight   Light
	yellowLight  Light
	redLight     Light
}

func NewTrafficLight() *TrafficLight {
	tl := &TrafficLight{}

	greenLight := &GreenLight{
		tl: tl,
	}

	yellowLight := &YellowLight{
		tl: tl,
	}

	redLight := &RedLight{
		tl: tl,
	}

	tl.setLight(greenLight)
	tl.greenLight = greenLight
	tl.yellowLight = yellowLight
	tl.redLight = redLight

	return tl
}

func (t *TrafficLight) GreenLight() error {
	return t.currentLight.GreenLight()
}

func (t *TrafficLight) YellowLight() error {
	return t.currentLight.YellowLight()
}

func (t *TrafficLight) RedLight() error {
	return t.currentLight.RedLight()
}

func (t *TrafficLight) setLight(l Light) {
	t.currentLight = l
}

func Countdown(rawSymbols string) {
	symbols := strings.Split(rawSymbols, " ")

	for _, symbol := range symbols {
		fmt.Printf("\r%s", symbol)
		time.Sleep(1 * time.Second)
	}
}
