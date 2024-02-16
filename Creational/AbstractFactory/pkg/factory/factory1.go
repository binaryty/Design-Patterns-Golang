package factory

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/controls"
)

type Factory1 struct {
}

func (n *Factory1) CreateButton(text string) controls.Button {
	return controls.NewButton1(text)
}

func (n *Factory1) CreateLabel(text string) controls.Label {
	return controls.NewLabel1(text)
}
