package factory

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/controls"
)

type ControlFactory interface {
	CreateButton(text string) controls.Button
	CreateLabel(text string) controls.Label
}
