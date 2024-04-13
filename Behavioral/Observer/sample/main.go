package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Behavioral/Observer/sample/pkg"
)

func main() {
	sub1 := &pkg.Subscriber{Name: "Sub-1"}
	sub2 := &pkg.Subscriber{Name: "Sub-2"}
	sub3 := &pkg.Subscriber{Name: "Sub-3"}
	sub4 := &pkg.Subscriber{Name: "Sub-4"}

	ch := pkg.Publisher{
		Name:      "Publisher channel",
		Consumers: map[string]pkg.Consumer{},
	}

	ch.Subscribe(sub1)
	ch.Subscribe(sub2)
	ch.Subscribe(sub3)
	ch.Subscribe(sub4)

	ch.Notify()

	fmt.Println("Unsubscribe Sub-3")
	ch.UnSubscribe(sub3)

	ch.Notify()
}
