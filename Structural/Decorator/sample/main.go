package main

import "fmt"

type InterfaceComponent interface {
	DoOperation()
}

type MainComponent struct {
}

func (c MainComponent) DoOperation() {
	fmt.Println("World!")
}

type Decorator struct {
	component InterfaceComponent
}

func NewDecorator(c InterfaceComponent) *Decorator {
	return &Decorator{
		component: c,
	}
}

func (d *Decorator) DoOperation() {
	d.component.DoOperation()
}

func (d *Decorator) NewOperation() {
	fmt.Println("Do Nothing")
}

type DecoratorSpace struct {
	*Decorator
}

func NewDecoratorSpace(c InterfaceComponent) *DecoratorSpace {
	return &DecoratorSpace{
		Decorator: NewDecorator(c),
	}
}

func (s *DecoratorSpace) DoOperation() {
	fmt.Print(" ")
	s.Decorator.DoOperation()
}

func (s *DecoratorSpace) NewOperation() {
	fmt.Println("New space operation")
}

type DecoratorComma struct {
	*Decorator
}

func NewDecoratorComma(c InterfaceComponent) *DecoratorComma {
	return &DecoratorComma{
		Decorator: NewDecorator(c),
	}
}

func (c *DecoratorComma) DoOperation() {
	fmt.Print(",")
	c.Decorator.DoOperation()
}

func (c *DecoratorComma) NewOperation() {
	fmt.Println("New comma operation")
}

type DecoratorHello struct {
	*Decorator
}

func NewDecoratorHello(c InterfaceComponent) *DecoratorHello {
	return &DecoratorHello{
		Decorator: NewDecorator(c),
	}
}

func (h *DecoratorHello) DoOperation() {
	fmt.Print("Hello")
	h.Decorator.DoOperation()
}

func (h *DecoratorHello) NewOperation() {
	fmt.Println("New hello operation")
}

func main() {
	component := NewDecoratorHello(
		NewDecoratorComma(
			NewDecoratorSpace(
				MainComponent{},
			),
		),
	)

	component.DoOperation()
	component.NewOperation()
}
