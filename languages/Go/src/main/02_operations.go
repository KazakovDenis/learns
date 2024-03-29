/* https://go.dev/ref/spec#Operators_and_punctuation
Сравнение: ==, !=, >, >=, <, <=,
Арифметика: +, -, /, *, %, +=, -=, *=, /=, %= [нет оператора остатка от деления], ++, --
Логические: && = and, || = or, ! = not
Побитовые: & = and, | = or, ~ = not (унарная), ^ - исключающее ИЛИ, а также &=, |=, ^=
Сдвиговые: >>, <<
*/
package main

var a, b int
var c bool

func main() {
	// инкремент/декремент только после переменной
	a++
	b &= ^a
	println(a, b, !c)
}
