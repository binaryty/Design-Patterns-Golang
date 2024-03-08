package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Behavioral/State/sample/pkg"
	"log"
)

func main() {
	vm := pkg.NewVendingMachine(1, 10)
	if err := vm.RequestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	if err := vm.AddItem(2); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	if err := vm.RequestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}
