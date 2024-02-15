package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/FactoryMethod/factory"
)

/*
Тестирование разработанной системы классов.

Даны пять строк. Используя конкретных создателей 1 и 2,
применить к каждой из данных строк метод AnOperation и вывести возвращаемый результат этого метода
(вначале выводятся результаты для первой строки, затем для второй и т. д.).
*/
func main() {
	var strings = []string{"Hello world!", "Test title", "Lore*Ipsum*Dolor", "abdcdefg", "g"}

	for _, s := range strings {
		product1 := factory.ConcreteCreator1{}
		product2 := factory.ConcreteCreator2{}

		fmt.Println(product1.AnOperation(s))
		fmt.Println(product2.AnOperation(s))
	}

}
