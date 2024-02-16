package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/Builder/pkg"
)

func main() {
	pascalBuilder := pkg.NewBuilderPascal()
	pythonBuilder := pkg.NewBuilderPython()
	cBuilder := pkg.NewBuilderC()

	pascalDirector := pkg.NewDirector(pascalBuilder)
	pascalDirector.Construct("string description for Pascal")
	fmt.Println(pascalDirector.GetResult())

	pythonDirector := pkg.NewDirector(pythonBuilder)
	pythonDirector.Construct("string description for Python")
	fmt.Println(pythonDirector.GetResult())

	cDirector := pkg.NewDirector(cBuilder)
	cDirector.Construct("string description for C")
	fmt.Println(cDirector.GetResult())
}
