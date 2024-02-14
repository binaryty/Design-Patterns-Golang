/*
Абстрактный класс Target содержит три абстрактных метода: GetA, GetB и Request
(не имеют параметров, возвращают значение целого типа).

Класс ConcreteTarget является потомком класса Target; он содержит поля A и B целого типа,
которые инициализируются в конструкторе, имеющем одноименные параметры.

Методы GetA и GetB класса ConcreteTarget возвращают значения полей A и B соответственно,
а метод Request возвращает сумму этих полей.

Класс Adaptee содержит два целочисленных поля A и B, конструктор с параметрами A и B,
задающий значения этих полей, метод GetAll, возвращающий текущие значения полей (либо с помощью выходных параметров,
либо с помощью возвращаемого значения — массива или кортежа),
и метод SpecificRequest без параметров, возвращающий произведение полей A и B.

Реализовать класс Adapter, адаптирующий класс Adaptee к интерфейсу класса Target.
Класс должен быть адаптером объекта: он порождается от класса Target
и включает ссылку ad на экземпляр адаптируемого объекта Adaptee.
Ссылка ad инициализируется в конструкторе класса Adapter путем вызова конструктора класса Adaptee с параметрами a и b,
совпадающими с одноименными параметрами конструктора класса Adapter.
Метод Request класса Adapter должен вызывать метод SpecificRequest объекта ad,
а методы GetA и GetB — возвращать значения полей a и b объекта ad, используя его метод GetAll
*/
package adapter

// Target интерфейс содержит три абстрактных метода: GetA, GetB и Request (не имеют параметров, возвращают значение целого типа).
type Target interface {
	GetA() int
	GetB() int
	Request() int
}

// ConcreteTarget является реализацией интефейса Target, он содержит поля A и B целого типа.
type ConcreteTarget struct {
	A int
	B int
}

// NewConcreteTarget конструктор ConcreteTarget с параметрами a и b, задающий значения этих полей.
func NewConcreteTarget(a int, b int) ConcreteTarget {
	return ConcreteTarget{
		A: a,
		B: b,
	}
}

// GetA возвращает значение поля A.
func (c ConcreteTarget) GetA() int {
	return c.A
}

// GetB возвращает значение поля B.
func (c ConcreteTarget) GetB() int {
	return c.B
}

// Request возвращает сумму полей A и B.
func (c ConcreteTarget) Request() int {
	return c.A + c.B
}

// Adaptee содержит два целочисленных поля A и B.
type Adaptee struct {
	A int
	B int
}

// NewAdaptee конструктор Adaptee с параметрами a и b, задающий значения этих полей.
func NewAdaptee(a int, b int) Adaptee {
	return Adaptee{
		A: a,
		B: b,
	}
}

// GetAll возвращающий текущие значения полей c помощью слайса.
func (a Adaptee) GetAll() []int {
	return []int{a.A, a.B}
}

// SpecificRequest без параметров, возвращающий произведение полей A и B.
func (a Adaptee) SpecificRequest() int {
	return a.A * a.B
}

// Adapter адаптирующий Adaptee к интерфейсу класса Target.
type Adapter struct {
	*Adaptee
}

// GetA возвращат значение поля a объекта Adaptee, используя его метод GetAll.
func (a Adapter) GetA() int {
	return a.GetAll()[0]
}

// GetB возвращат значение поля b объекта Adaptee, используя его метод GetAll.
func (a Adapter) GetB() int {
	return a.GetAll()[1]
}

// Request метод Adapter должен вызывать метод SpecificRequest объекта Adaptee.
func (a Adapter) Request() int {
	return a.SpecificRequest()
}
