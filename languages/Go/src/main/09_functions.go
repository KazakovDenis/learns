package main

import (
	"fmt"
	"time"
)

// Объявление функции
// Нельзя декларировать функцию внутри другой функции
// Аргументы передаются функции путём копирования значения
// func Function(arg1 arg1type, arg2 arg2type) returnType {
//     тело функции
// }

// Но можно объявить внутри другой с помощью литеральной
// формы синтаксиса (по сути - лямбда-функция)
func foo() {
	f := func(s string) string { return s }
	fmt.Println("Literal:", f)
}

// Результат функции тоже можно именовать
// Эти переменные будут проинициализированы по умолчанию
// тогда в инструкции return имя можно не указывать
func divide(x int) (half int) {
	half = x / 2
	return // результат всё равно будет возвращён
}

// Параметры одного типа по порядку можно записать через запятую
func sum(x, y int) int {
	return x + y
}

// Функция с переменным кол-вом аргументов (variadic)
// Такой параметр должен следовать последним
// Внутри функции он рассматривается как слайс
func multiply(x ...int) int {
	result := 0
	for item := range x {
		result += item
	}
	return result
}

// Может возвращать более 1 значения
// Если значения не нужны: _, x, _ := ReturnsMany()
func returnsMany() (int, int, int) {
	return 1, 2, 3
}

// Допускаются рекурсивные вызовы
func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

// Функция первого класса
// ======================
// Тип функции виден в её сигнатуре, то есть определяется как
// набор типов и количества аргументов, возвращаемых значений.
// Эта функция имеет тип func(string) string
func echo(str string) (v string) {
	return str
}

// Её можно присвоить другой переменной такого же типа
func renameEcho() {
	var rEcho func(string) string
	rEcho = echo
	fmt.Println("Renamed echo:", rEcho)
}

// Функцию можно передавать в другую функцию следующим образом
func do(who string, how func(string) string) {
	fmt.Println(how(who))
}

// Замыкания
// =========
// Go — язык с лексической областью видимости (lexically scoped).
// Это значит, что переменные, определённые в окружающих блоках видимости (например,
// глобальные переменные), доступны функции всегда, а не только на время вызова.
func generate(seed int) func() {
	return func() {
		// замыкание получает внешнюю переменную seed
		fmt.Println("Generated:", seed)
		seed += 1
	}
}

// Пример декоратора для подсчета времени выполнения
func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start))
	}
}

// Специальная функция init()
// В пакете и даже в одном файле можно декларировать несколько таких функций.
// Они будут вызваны один раз при инициализации пакета, после присвоения
// глобальных переменных, в том порядке, в котором они предоставлены компилятору.
func init() {}

// funcopts - подход для инициализации объектов
// tree := newTree(Option1(10), Option2(20))
func newTree(opts ...Option) *Tree {
	i := &Tree{}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

type Option func(*Tree)

func Option1(option1 int) Option {
	return func(i *Tree) {
		i.Value = option1
	}
}

func main() {
	renameEcho()
	do("hello", echo)
	generator := generate(0)
	generator()
	generator()
}
