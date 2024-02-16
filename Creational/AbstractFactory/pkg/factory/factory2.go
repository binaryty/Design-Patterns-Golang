package factory

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/controls"
)

type Factory2 struct {
}

func (y *Factory2) CreateButton(text string) controls.Button {
	return controls.NewButton2(text)
}

func (y *Factory2) CreateLabel(text string) controls.Label {
	return controls.NewLabel2(text)
}
