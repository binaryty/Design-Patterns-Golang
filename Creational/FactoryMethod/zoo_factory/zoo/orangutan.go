package zoo

import (
	"fmt"
	"reflect"
)

type Orangutan struct {
	Name string
}

func NewOrangutan(name string) *Orangutan {
	return &Orangutan{
		Name: name,
	}
}

func (o Orangutan) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(o).Name(), o.Name)
}
