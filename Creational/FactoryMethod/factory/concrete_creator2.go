package factory

type ConcreteCreator2 struct {
}

// FactoryMethod принимает на вход строковый параметр info, возвращающий ссылку на объект Product
func (c ConcreteCreator2) FactoryMethod(info string) Product {
	return NewConcreteProduct2(info)
}

// AnOperation создает продукт с помощью фабричного метода, передавая ему параметр info,
// дважды вызывает метод Transform созданного продукта и с помощью его метода GetInfo возвращает полученный результат.
func (c ConcreteCreator2) AnOperation(info string) string {
	product := c.FactoryMethod(info)

	product.Transform()
	product.Transform()

	return product.GetInfo()
}
