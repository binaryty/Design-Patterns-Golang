package controls

import (
	"fmt"
	"strings"
)

type Label2 struct {
	text string
}

func NewLabel2(text string) *Label2 {
	return &Label2{
		text: text,
	}
}

func (l Label2) GetControl() string {
	return fmt.Sprintf("=%s=", strings.ToUpper(l.text))
}
