package controls

import (
	"fmt"
	"strings"
)

type Label1 struct {
	text string
}

func NewLabel1(text string) *Label1 {
	return &Label1{
		text: text,
	}
}

func (l Label1) GetControl() string {
	return fmt.Sprintf("\"%s\"", strings.ToLower(l.text))
}
