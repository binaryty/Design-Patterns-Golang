package zoo

import (
	"fmt"
	"reflect"
)

type Gorilla struct {
	Name string
}

func NewGorilla(name string) *Gorilla {
	return &Gorilla{
		Name: name,
	}
}

func (g Gorilla) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(g).Name(), g.Name)
}
