package pkg

import "strings"

type BuilderC struct {
	product string
}

func NewBuilderC() *BuilderC {
	return &BuilderC{product: ""}
}

func (b *BuilderC) BuildStart(char rune) {
	b.product = b.product + strings.ToLower(string(char))
}

func (b *BuilderC) BuildFirstChar(char rune) {
	b.product = b.product + strings.ToLower(string(char))
}

func (b *BuilderC) BuildNextChar(char rune) {
	b.product = b.product + strings.ToLower(string(char))
}

func (b *BuilderC) BuildDelim() {
}

func (b *BuilderC) GetResult() string {
	return b.product
}
