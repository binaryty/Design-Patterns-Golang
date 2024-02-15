package zoo

import (
	"fmt"
	"reflect"
)

type Lion struct {
	Name string
}

func NewLion(name string) *Lion {
	return &Lion{
		Name: name,
	}
}

func (l Lion) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(l).Name(), l.Name)
}
