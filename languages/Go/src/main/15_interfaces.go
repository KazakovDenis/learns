package main

import (
	"fmt"
	"io"
)

// Worker
// Интерфейс — это набор методов, которые могут быть реализованы типом.
// Имена аргументов могут не указываться, но это ненаглядно:
// Do(int) int
type Worker interface {
	Work(task string)
	Retire()
}

type Company struct {
	personal []Worker
}

// Hire
// newbie может быть любого типа, удовлетворяющему интерфейсу Worker
// Удовлетворять интерфейсу == реализовывать все его методы
func (c *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
	newbie.Work("development")
}

// Man - структура, удовлетворяющая интерфейсу Worker
type Man struct {
	name string
}

func (m Man) Work(task string) {
	fmt.Printf("%s did the task: %s\n", m.name, task)
}

// Методы принимают Man по-разному: по значению и по указателю.
// Поэтому если передать по значению, код не скомплириуется,
// т.к. указатель на тип уже не реализовывает интерфейс.
// Поэтому передавать в Hire будем указатель на тип,

func (m *Man) Retire() {
	fmt.Printf("%s retired\n", m.name)
}

// Явная проверка реализации типом Man интерфейса Worker
var _ Worker = (*Man)(nil)

/*
Интерфейсы должны быть компактными, даже из одного метода.
В таком случае они могут именоваться с суффиксом -er, e.g.: Reader.
Интерфейсы можно встраивать в другие интерфейсы.
Интерфейс описывается в том же пакете, в котором применяется.

Стандартная библиотека содержит различные интерфейсы, например:
Stringer, Reader, Writer.
Вспомогательные функции для Reader / Writer:
io.Copy - скопировать из Reader в Writer
io.CopyN - скопировать N байт из Reader в Writer
io.ReadAll - прочитать всё из Reader
io.ReadAtLeast - прочитать минимум байт из Reader, иначе io.ErrUnexpectedEOF
*/

// Hasher - расширяет интерфейс стандартной библиотеки Writer методом Hash()
type Hasher interface {
	io.Writer
	Hash() byte
}

// PassAnyType
// Чтобы функция могла принимать любой параметр, можно
// указать пустой интерфейс, any - синоним к пустому интерфейсу
//func PassAnyType(arg interface{}) {}
func PassAnyType(interfaceArg any) {
	// Однако с этим аргументом ничего нельзя сделать
	// до приведения интерфейса к типу: interfaceArg.(type)
	// Интерфейсная переменная относится к ссылочному типу.
	// Интерфейс под капотом представляет собой 2 указателя:
	// на обёрнутую переменную и на информацию о её типе.
	// Преобразование типа изменяет информацию о типе либо возвращает обёрнутую переменную.
	// Интерфейсы сравниваются по цепочке: сначала тип, затем данные.
	i := interfaceArg.(int)     // если v не число, то будет паника
	j, ok := interfaceArg.(int) // если v не число, паники не будет, ok=false
	fmt.Printf("Empty interface: %d, %v, %t \n", i, j, ok)
}

type EmptyStruct struct{}

func IsNil(obj interface{}) bool {
	// Переданная структура будет обернута интерфейсом, то есть
	// если передать пустую структуру, то она будет указывать
	// на nil-значение, но не будет равна ему
	return obj == nil
}

func main() {
	company := Company{}
	worker := Man{"Rob"}
	company.Hire(&worker)
	worker.Retire()
	PassAnyType(1)

	s := EmptyStruct{}
	fmt.Println("IsNil:", IsNil(s))
}
