package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/factory"
	"strings"
)

func main() {
	set := []string{"B Accept", "L Edit1", "B Ok", "L Name", "B Yes", "L Address"}

	f1 := &factory.Factory2{}
	Client1 := pkg.NewClient(f1)

	f2 := &factory.Factory1{}
	Client2 := pkg.NewClient(f2)

	for _, s := range set {
		params := strings.Split(s, " ")

		ctrlType := params[0]
		text := params[1]

		switch ctrlType {
		case "B":
			Client1.AddButton(text)
			Client2.AddButton(text)
		case "L":
			Client1.AddLabel(text)
			Client2.AddLabel(text)
		}
	}

	fmt.Println(Client1.GetControls())
	fmt.Println(Client2.GetControls())
}
