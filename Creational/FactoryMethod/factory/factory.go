package factory

// Creator интерфейс создателя.
type Creator interface {
	FactoryMethod(info string) Product
	AnOperation(info string) string
}

// Product интерфейс, который содержит два абстрактных метода, связанных с получением и преобразованием строки.
type Product interface {
	GetInfo() string
	Transform()
}
