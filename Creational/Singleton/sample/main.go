package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/Singleton/sample/pkg"
	"strconv"
)

var singleton *pkg.Singleton

func init() {
	fmt.Println("initialization")
	singleton = &pkg.Singleton{"Одиночка"}
}

func main() {
	for i := 0; i < 3; i++ {
		pkg.NewSingleton(singleton, "Create singleton "+strconv.Itoa(i))
	}
}
