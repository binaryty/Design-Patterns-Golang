package pkg

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/factory"
	"strings"
)

type Client struct {
	f        factory.ControlFactory
	controls []string
}

func NewClient(f factory.ControlFactory) *Client {
	return &Client{
		f:        f,
		controls: make([]string, 0),
	}
}

func (c *Client) AddButton(text string) {
	button := c.f.CreateButton(text)
	c.controls = append(c.controls, button.GetControl())
}

func (c *Client) AddLabel(text string) {
	label := c.f.CreateLabel(text)
	c.controls = append(c.controls, label.GetControl())
}

func (c *Client) GetControls() string {
	return strings.Join(c.controls, " ")
}
