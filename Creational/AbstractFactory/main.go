package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg"
	"github.com/binaryty/Design-Design-Patterns-Golang/Creational/AbstractFactory/pkg/factory"
	"strings"
)

/*
Дано целое число Nf, которое может быть равно 1 или 2, целые числа Na и Nb и строка S, содержащая символы «A» и «B».
Описать ссылочные переменные f типа AbstractFactory, pa типа AbstractProductA и pb типа AbstractProductB.
Если число Nf равно 1, то связать f с конкретной фабрикой 1, если Nf равно 2, то связать f с конкретной фабрикой 2.
Используя фабричные методы созданной фабрики, создать конкретные продукты типа A и B,
инициализировав их данными числами Na и Nb соответственно, и связать их с переменными pa и pb.
Затем для созданных продуктов pa и pb выполнить методы A и B в порядке, указанном в исходной строке S.
При этом метод A должен вызываться для продукта pa, а метод B — для продукта pb,
причем параметром метода B должен быть продукт pa.
Используя методы GetInfo, вывести итоговые значения объектов-продуктов pa и pb (в указанном порядке).

Примечание. При выполнении задания используются только ссылки на абстрактные классы,
а также только методы, определенные в абстрактных классах (за исключением конструктора создаваемой конкретной фабрики),
что и составляет суть паттерна «Абстрактная фабрика».
*/
func main() {
	const N = 6

	var set [N]string = [N]string{"B Accept", "L Edit1", "B Ok", "L Name", "B Yes", "L Address"}

	f1 := &factory.Factory2{}
	Client1 := pkg.NewClient(f1)

	f2 := &factory.Factory1{}
	Client2 := pkg.NewClient(f2)

	for _, s := range set {
		params := strings.Split(s, " ")

		ctrlType := params[0]
		text := params[1]

		switch ctrlType {
		case "B":
			Client1.AddButton(text)
			Client2.AddButton(text)
		case "L":
			Client1.AddLabel(text)
			Client2.AddLabel(text)
		}
	}

	fmt.Println(Client1.GetControls())
	fmt.Println(Client2.GetControls())
}
