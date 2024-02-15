package factory

import "strings"

type ConcreteProduct1 struct {
	Info string
}

// NewConcreteProduct1 создает экземпляр ConcreteProduct1 инициализирует поле info и преобразует его к нижнему регистру.
func NewConcreteProduct1(info string) *ConcreteProduct1 {
	return &ConcreteProduct1{
		Info: strings.ToLower(info),
	}
}

// GetInfo возвращает текущее значение поля info.
func (c *ConcreteProduct1) GetInfo() string {
	return c.Info
}

// Transform добавляет дополнительный пробел после каждого непробельного символа поля info(кроме его последнего символа).
func (c *ConcreteProduct1) Transform() {
	var sb strings.Builder

	l := len(c.Info)

	for i := 0; i < l-1; i++ {
		if c.Info[i] != ' ' {
			sb.WriteByte(c.Info[i])
			sb.WriteByte(' ')
		} else {
			sb.WriteByte(c.Info[i])
		}
	}

	sb.WriteByte(c.Info[l-1])

	c.Info = sb.String()
}
