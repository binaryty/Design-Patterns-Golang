package controls

import (
	"fmt"
	"strings"
)

type Button2 struct {
	text string
}

func NewButton2(text string) *Button2 {
	return &Button2{
		text: text,
	}
}

func (b Button2) GetControl() string {
	return fmt.Sprintf("[%s]", strings.ToUpper(b.text))
}
