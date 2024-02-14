/*
Тестирование разработанной системы адаптера.
Дано целое число N (≤ 6) и набор из N троек (a, b, c),
где c является символом «+» или «*», а элементы a и b являются целыми числами.
Создать структуру данных (например, массив) с элементами-ссылками типа Target
и заполнить ее объектами типа ConcreteTarget (для троек с символом «+»)
и Adapter (для троек с символом «*») с полями, равными a и b.
Перебирая элементы полученного набора в обратном порядке,
вывести для каждого из них значения полей a, b и результат выполнения метода Request.
*/
package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Adapter/adapter"
)

const N = 3

func main() {

	var triple = []struct {
		a int
		b int
		c byte
	}{
		{
			a: 1,
			b: 2,
			c: '+',
		},
		{
			a: 3,
			b: 2,
			c: '*',
		},
		{
			a: 5,
			b: 26,
			c: '+',
		},
	}

	targets := make([]adapter.Target, N)

	for i := 0; i < N; i++ {
		t := triple[i]
		if t.c == '+' {
			targets[i] = &adapter.ConcreteTarget{
				A: t.a,
				B: t.b,
			}
		} else if t.c == '*' {
			targets[i] = adapter.Adapter{&adapter.Adaptee{
				A: t.a,
				B: t.b,
			}}
		}
	}

	for i := N - 1; i >= 0; i-- {
		target := targets[i]
		fmt.Printf("a: %d\n", target.GetA())
		fmt.Printf("b: %d\n", target.GetB())
		fmt.Printf("Request: %d\n", target.Request())
	}
}
