package zoo

import (
	"fmt"
	"reflect"
)

type Leopard struct {
	Name string
}

func NewLeopard(name string) *Leopard {
	return &Leopard{
		Name: name,
	}
}

func (l Leopard) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(l).Name(), l.Name)
}
