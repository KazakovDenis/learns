// https://go.dev/ref/spec#Types
// https://go.dev/ref/spec#Integer_literals
// http://golang-book.ru/chapter-03-types.html
package main

import "fmt"

// Длинное объявление
var str1 string
var str2 string = "string2"

// untyped float const
const pi = 3.14159

const (
	// typed const
	flag     uint8 = 128
	name           = "Denis Kazakov"
	fullName       // равно предыдущей константе
)

func main() {
	str1 = "string1"
	println(str1)

	// Константы
	println("pi", pi, "flag", flag, "name", name, "fullName", fullName)
	var (
		// нетипизированные константы могут быть присвоены переменным других типов
		a float32 = pi
		b float64 = pi
	)
	println("a", a, "b", b)

	// Без указания типа, определяется при первом присвоении
	var text = "text"
	println(text)

	// Краткая нотация
	integer := 10
	fmt.Printf("Type of %d is %T\n", integer, integer)

	// Короткая нотация с приведением к недефолтному типу
	// Эквивалент: var floatVar float32 = 101.3
	float32Var := float32(101.3)
	fmt.Printf("Short notation with a cast: %f\n", float32Var)

	// Обращение по индексу возвращает номер по ASCII
	// Приведение к строке преобразует обратно к символу
	var symbol = string(str2[0])
	println("First symbol:", symbol)

	// Создание алиаса типа
	type float = float32
	// Переменная инициализируется пустым значением своего типа
	// Можно указывать более одной через запятую
	var floatingPoint, oneMore float
	println("Empty floats:", floatingPoint, oneMore)

	var (
		height, length int
		weight         float64
		char           byte = 'G'
		language            = "Go"
	)
	println(height, length, weight, char, language)

	// Создание типа
	type Name string
	var name Name = "Denis"
	println("My name is", name)

	// Создание перечисления, iota позволяет инкрементировать
	// последующие значения в блоке, также можно указать множитель
	const (
		Black = iota * 10
		Gray
		White
	)
	println("Colours", Black, Gray, White)

	type Weekday int
	// iota изменяется с каждой строкой
	const (
		Monday Weekday = iota + 1
		Tuesday
	)
	print("IntEnum", Monday, Tuesday)
}
