package zoo_factory

import (
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/FactoryMethod/zoo_factory/zoo"
)

const (
	LionID = iota
	TigerID
	LeopardID
	GorillaID
	OrangutanID
	ChimpanzeeID
)

type Animal interface {
	GetInfo() string
}

type AnimalCreator interface {
	GetZoo(idxes []int, names []string) []Animal
	CreateAnimal(ind int, name string) Animal
}

type CatCreator struct {
}

type ApeCreator struct {
}

// CreateAnimal принимает параметр целого типа ind и строковый параметр name и возвращает объект типа Animal.
// параметр ind метода CreateAnimal определяет тип создаваемого животного
// по его индексу в группе кошачьих, а параметр name — имя животного.
func (c CatCreator) CreateAnimal(ind int, name string) Animal {
	var catAnimal Animal
	switch ind {
	case LionID:
		catAnimal = zoo.NewLion(name)
	case TigerID:
		catAnimal = zoo.NewTiger(name)
	case LeopardID:
		catAnimal = zoo.NewLeopard(name)
	default:
		panic("unknown cat animal")
	}

	return catAnimal
}

func (c CatCreator) GetZoo(idxes []int, names []string) []Animal {
	animals := make([]Animal, len(idxes))

	for i, v := range idxes {
		animals[i] = c.CreateAnimal(v, names[i])
	}

	return animals
}

// CreateAnimal принимает параметр целого типа ind и строковый параметр name и возвращает объект типа Animal.
// параметр ind метода CreateAnimal определяет тип создаваемого животного
// по его индексу в группе обезьян, а параметр name — имя животного.
func (c ApeCreator) CreateAnimal(ind int, name string) Animal {
	var apeAnimal Animal

	switch ind {
	case GorillaID:
		apeAnimal = zoo.NewGorilla(name)
	case OrangutanID:
		apeAnimal = zoo.NewOrangutan(name)
	case ChimpanzeeID:
		apeAnimal = zoo.NewChimpanzee(name)
	default:
		panic("unknown ape animal")
	}

	return apeAnimal
}

func (c ApeCreator) GetZoo(idxes []int, names []string) []Animal {
	animals := make([]Animal, len(idxes))

	for i, v := range idxes {
		animals[i] = c.CreateAnimal(v, names[i])
	}

	return animals
}
