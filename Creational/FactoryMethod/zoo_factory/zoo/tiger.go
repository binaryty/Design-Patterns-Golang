package zoo

import (
	"fmt"
	"reflect"
)

type Tiger struct {
	Name string
}

func NewTiger(name string) *Tiger {
	return &Tiger{
		Name: name,
	}
}

func (t Tiger) GetInfo() string {
	return fmt.Sprintf("%s %s", reflect.TypeOf(t).Name(), t.Name)
}
