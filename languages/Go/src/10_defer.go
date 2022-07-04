package main

import "fmt"

// defer - оператор отложенного вызова, используют аналогично
// контекстным менеджерам в питоне - для управления ресурсами.
// defer вычисляет аргументы для вызова, но вызов выполняется
// непосредственно перед тем, как отложившая его функция вернёт управление.
// Функция может содержать более 1 оператора defer, все последующие
// складываются в стек и выполняются после return в обратном порядке
func deferredExec() {
	defer fmt.Println("deferred func")
	fmt.Println("evaluated func")
}

// defer работает и с литералом (лямбдой) функции
// и может использовать замыкание этой функции
func deferredLambda() {
	defer func() { fmt.Println("deferred lambda") }()
	fmt.Println("evaluated lambda")
}

// аргументы вычисляются во время вызова defer, а не отложенной функции
func deferOperand() {
	a := "some text"
	defer func(s string) { fmt.Println(s) }(a)
	a = "another text"
}

// отложенные функции могут перехватывать панику,
// что похоже на try-except в Питоне
func catchPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic is caught:", r)
		}
	}()
	panic("Unexpected error!")
}

func main() {
	deferredExec()
	deferredLambda()
	deferOperand()
	catchPanic()
}
