/*
Абстрактный класс Target содержит три абстрактных метода: GetA, GetB и Request
(не имеют параметров, возвращают значение целого типа).

Класс ConcreteTarget является потомком класса Target; он содержит поля a и b целого типа,
которые инициализируются в конструкторе, имеющем одноименные параметры.

Методы GetA и GetB класса ConcreteTarget возвращают значения полей a и b соответственно,
а метод Request возвращает сумму этих полей.

Класс Adaptee содержит два целочисленных поля a и b, конструктор с параметрами a и b,
задающий значения этих полей, метод GetAll, возвращающий текущие значения полей (либо с помощью выходных параметров,
либо с помощью возвращаемого значения — массива или кортежа),
и метод SpecificRequest без параметров, возвращающий произведение полей a и b.

Реализовать класс Adapter, адаптирующий класс Adaptee к интерфейсу класса Target.
Класс должен быть адаптером объекта: он порождается от класса Target
и включает ссылку ad на экземпляр адаптируемого объекта Adaptee.
Ссылка ad инициализируется в конструкторе класса Adapter путем вызова конструктора класса Adaptee с параметрами a и b,
совпадающими с одноименными параметрами конструктора класса Adapter.
Метод Request класса Adapter должен вызывать метод SpecificRequest объекта ad,
а методы GetA и GetB — возвращать значения полей a и b объекта ad, используя его метод GetAll
*/
package main

// Target интерфейс содержит три абстрактных метода: GetA, GetB и Request (не имеют параметров, возвращают значение целого типа).
type Target interface {
	GetA() int
	GetB() int
	Request() int
}

// ConcreteTarget является реализацией интефейса Target, он содержит поля a и b целого типа.
type ConcreteTarget struct {
	a int
	b int
}

// NewConcreteTarget конструктор ConcreteTarget с параметрами a и b, задающий значения этих полей.
func NewConcreteTarget(a int, b int) ConcreteTarget {
	return ConcreteTarget{
		a: a,
		b: b,
	}
}

// GetA возвращает значение поля a.
func (c ConcreteTarget) GetA() int {
	return c.a
}

// GetB возвращает значение поля b.
func (c ConcreteTarget) GetB() int {
	return c.b
}

// Request возвращает сумму полей a и b.
func (c ConcreteTarget) Request() int {
	return c.a + c.b
}

// Adaptee содержит два целочисленных поля a и b.
type Adaptee struct {
	a int
	b int
}

// NewAdaptee конструктор Adaptee с параметрами a и b, задающий значения этих полей.
func NewAdaptee(a int, b int) Adaptee {
	return Adaptee{
		a: a,
		b: b,
	}
}

// GetAll возвращающий текущие значения полей c помощью слайса.
func (a Adaptee) GetAll() []int {
	return []int{a.a, a.b}
}

// SpecificRequest без параметров, возвращающий произведение полей a и b.
func (a Adaptee) SpecificRequest() int {
	return a.a * a.b
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
