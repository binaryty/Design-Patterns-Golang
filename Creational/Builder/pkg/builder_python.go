package pkg

import "strings"

type BuilderPython struct {
	product string
}

func NewBuilderPython() *BuilderPython {
	return &BuilderPython{product: ""}
}

func (b *BuilderPython) BuildStart(char rune) {
	b.product += strings.ToLower(string(char))
}

func (b *BuilderPython) BuildFirstChar(char rune) {
	b.product += strings.ToLower(string(char))
}

func (b *BuilderPython) BuildNextChar(char rune) {
	b.product += strings.ToLower(string(char))
}

func (b *BuilderPython) BuildDelim() {
	b.product += "_"
}

func (b *BuilderPython) GetResult() string {
	return b.product
}
