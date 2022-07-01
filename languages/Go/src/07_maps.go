// https://go.dev/ref/spec#Map_types
package main

import "fmt"

func main() {
	// В пустую объявленную мапу присвоить значения нельзя
	// var m map[KeyType]ValueType == nil

	// Map требуется инициализировать через make
	// make - универсальный контруктор всех ссылочных типов
	map1 := make(map[string]string)
	map1["foo"] = "foo"
	fmt.Println("Create map:", map1)

	// Две переменные ссылочного типа могут указывать на один и тот же объект
	map2 := map1
	map2["bar"] = "bar"
	fmt.Println("Change map via ref:", map1)

	// Создание через composite literal
	map3 := map[string]int{"first": 1, "second": 2}
	fmt.Println("Composite literal:", map3)

	// Для ключей должны быть определены операторы == и !=,
	// поэтому ключ не может быть функцией, хеш-таблицей или слайсом
	// var MyMap map[[]byte]string - ошибка компиляции
	// На тип значений не накладывается никаких ограничений
	map4 := map[int][]any{1: {'a', 2}, 2: {"key"}}
	fmt.Println("Value type:", map4)

	var val1 = map4[1]
	val2, ok := map4[2]
	fmt.Println("Key, value:", val1, val2, ok)

	// ok возвращает нулевое значение типа, если ключ не задан,
	// поэтому его используют для проверки
	map5 := make(map[int]int)
	v1, ok1 := map5[100]
	map5[100] = 100
	v2, ok2 := map5[100]
	fmt.Println("Ok:", v1, ok1, v2, ok2)

	// Взять ссылку на элемент map нельзя
	// addr := &m[x] - ошибка компиляции

	// Получение кол-ва ключей
	var map6 map[int]int
	map7 := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println("Length", len(map6), len(map7))

	// Функция len() не даёт гарантии, что map инициализирована
	// для этого сравнивают с nil (единственное значение, с которым можно сравнивать map)
	var map8 map[string]string
	if map8 != nil { // если не проверить это условие,
		map8["foo"] = "bar" // то здесь можно получить panic
	}

	// Удаление
	delete(map8, "foo")
	fmt.Println("Delete:", map8)

	// Итерация по map не упорядочена
	for key, value := range map7 {
		// k будет перебирать ключи,
		// v — соответствующие этим ключам значения
		fmt.Printf("Key %v has value %v \n", key, value)
	}

	// Для безопасного использования map из нескольких потоков должны применяться
	// механизмы защиты, иначе возможно повреждение данных или состояние гонки
	// Для этого используется https://pkg.go.dev/sync#Map
}
