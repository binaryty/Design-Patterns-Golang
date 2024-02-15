package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/FactoryMethod/factory"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/FactoryMethod/zoo_factory"
)

func main() {
	/*
		Тестирование разработанной системы классов.

		Даны пять строк. Используя конкретных создателей 1 и 2,
		применить к каждой из данных строк метод AnOperation и вывести возвращаемый результат этого метода
		(вначале выводятся результаты для первой строки, затем для второй и т. д.).
	*/
	var strings = []string{"Hello world!", "Test title", "Lore*Ipsum*Dolor", "In trance we trust", "god is a dj"}

	for _, s := range strings {
		product1 := factory.ConcreteCreator1{}
		product2 := factory.ConcreteCreator2{}

		fmt.Println(product1.AnOperation(s))
		fmt.Println(product2.AnOperation(s))
	}

	/*
		Тестирование разработанной системы классов.

		Дан набор из 4 пар (ind, name), где ind — целое число в диапазоне от 0 до 2, а name — строка.
		Сформировать массивы idxes и names размера 4, содержащие числа и строки из исходного набора,
		и использовать их в методе GetZoo для создателей CatCreator и ApeCreator,
		получив в результате наборы кошачьих и обезьян размера 4.
		С помощью метода GetInfo вывести информацию о животных из каждого набора.
	*/
	var pairs = []struct {
		ind  int
		name string
	}{
		{zoo_factory.LionID, "Alex"},
		{zoo_factory.TigerID, "Frank"},
		{zoo_factory.LeopardID, "Bonny"},
		{zoo_factory.GorillaID, "Garry"},
		{zoo_factory.OrangutanID, "Rudi"},
		{zoo_factory.ChimpanzeeID, "Maria"},
	}

	idxes := make([]int, len(pairs))
	names := make([]string, len(pairs))

	for i, p := range pairs {
		idxes[i] = p.ind
		names[i] = p.name
	}

	catFactory := zoo_factory.CatCreator{}
	apeFactory := zoo_factory.ApeCreator{}

	catZoo := catFactory.GetZoo(idxes[:3], names[:3])
	apeZoo := apeFactory.GetZoo(idxes[3:], names[3:])

	for _, z := range catZoo {
		fmt.Println(z.GetInfo())
	}

	for _, z := range apeZoo {
		fmt.Println(z.GetInfo())
	}
}
