package pkg

import "strings"

type BuilderPascal struct {
	product string
}

func NewBuilderPascal() *BuilderPascal {
	return &BuilderPascal{product: ""}
}

func (b *BuilderPascal) BuildStart(ch rune) {
	b.product += strings.ToLower(string(ch))
}

func (b *BuilderPascal) BuildFirstChar(ch rune) {
	b.product += strings.ToUpper(string(ch))
}

func (b *BuilderPascal) BuildNextChar(ch rune) {
	b.product += strings.ToLower(string(ch))
}

func (b *BuilderPascal) BuildDelim() {

}

func (b *BuilderPascal) GetResult() string {
	return b.product
}
