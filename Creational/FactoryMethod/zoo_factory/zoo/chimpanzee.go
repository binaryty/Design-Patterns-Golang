package zoo

import (
	"fmt"
	"reflect"
)

type Chimpanzee struct {
	Name string
}

func NewChimpanzee(name string) *Chimpanzee {
	return &Chimpanzee{
		Name: name,
	}
}

func (c Chimpanzee) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(c).Name(), c.Name)
}
