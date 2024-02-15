package factory

import "strings"

type ConcreteProduct2 struct {
	Info string
}

// NewConcreteProduct2 создает экземпляр ConcreteProduct1 инициализирует поле info и преобразует его к верхнему регистру.
func NewConcreteProduct2(info string) *ConcreteProduct2 {
	return &ConcreteProduct2{
		Info: strings.ToUpper(info),
	}
}

// GetInfo возвращает текущее значение поля info.
func (c *ConcreteProduct2) GetInfo() string {
	return c.Info
}

// Transform добавляет два дополнительных символа * (звездочка) после каждого символа,отличного от звездочки (кроме последнего символа).
func (c *ConcreteProduct2) Transform() {
	var sb strings.Builder

	l := len(c.Info)

	for i := 0; i < l-1; i++ {
		if c.Info[i] != '*' {
			sb.WriteByte(c.Info[i])
			sb.WriteString("**")
		} else {
			sb.WriteByte(c.Info[i])
		}
	}
	sb.WriteByte(c.Info[l-1])

	c.Info = sb.String()
}
