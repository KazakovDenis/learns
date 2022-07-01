package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// Tree
// Поля структуры могут быть любыми типами
// В том числе указателями на саму структуру
type Tree struct {
	// Экспортируемое поле
	Value int
	// Неэкспортируемые поля
	leftChild  *Tree
	rightChild *Tree
}

// Request
// У каждого поля структуры может быть набор аннотаций (tags)
// Теги не влияют на представление или работу с данными напрямую, но могут
// использоваться пакетами для получения дополнительной информации о конкретном поле.
// Можно вводить свои теги и работать с ними через пакет reflect стандартной библиотеки.
type Request struct {
	StatusCode      int    `json:"status_code" yaml:"status_code"`
	Body            string `json:"body" yaml:"body" example:"CREATED"`
	NonSerializable int    `json:"-"` // - означает, что это поле не будет сериализовано
}

// Анонимная структура объявляется и сразу инициализируется
var anon = struct {
	NameContains string `json:"name_contains"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
}{
	NameContains: "Rob",
	Limit:        50,
}

func dumpRequest(code int, body string) {
	request := Request{
		StatusCode: code,
		Body:       body,
	}
	dump, err := json.Marshal(request)
	if err != nil {
		fmt.Println("unable marshal to json")
	}
	fmt.Printf("Serialized: %s \n", dump)
}

// NewTree
// Часто применяют конструкторы, пишут с префиксом New
// Если конструктор производит валидацию аргументов,
// функция должна возвращать ошибку последним аргументом
// Конструкторы используют также для инкапсуляции
// (экспортируютнеэкспортируемые структуры)
func NewTree(value int, left, right *Tree) (Tree, error) {
	return Tree{
		Value:      value,
		leftChild:  left,
		rightChild: right,
	}, nil
}

func main() {
	// Инициализация со значениями по умолчанию
	// var p Person
	tree1 := Tree{}
	fmt.Println("Init default tree:", tree1)

	// Неявное, с указанием всех значений по порядку
	tree2 := Tree{1, &tree1, &tree1}
	fmt.Println("Init implicit tree:", tree2)

	// Явное, с указанием названий полей, можно указывать не все
	tree3 := Tree{
		Value:      1,
		rightChild: &tree1,
	}
	fmt.Println("Init explicit tree:", tree3)

	tree4, _ := NewTree(3, &tree2, &tree3)
	fmt.Println("Constructor:", tree4)

	// Доступ к полям
	fmt.Println("Dot notation:", tree4.Value)

	// Сериализация
	dumpRequest(200, "OK")

	// Размер struct{} равен 0, при этом объект c имеет адрес
	fmt.Printf("Anon: %v \n", anon)
	c := struct{}{}
	fmt.Println("Empty anon size", unsafe.Sizeof(c))
	fmt.Println("Empty anon addr", unsafe.Pointer(&c))
}
