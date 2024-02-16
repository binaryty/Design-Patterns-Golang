package controls

import (
	"fmt"
	"strings"
)

type Button1 struct {
	text string
}

func NewButton1(text string) *Button1 {
	return &Button1{
		text: text,
	}
}

func (b Button1) GetControl() string {
	return fmt.Sprintf("<%s>", strings.ToLower(b.text))
}
