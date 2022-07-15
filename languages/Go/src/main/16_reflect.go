// https://pkg.go.dev/reflect
package main

/*
Рефлексия - возможность получить информацию о типе из переменной этого типа.
С помощью пакета reflect можно динамически создавать типы в рантайме.
*/

import (
	"fmt"
	"reflect"
)

type SomeType struct {
	IntField int
	StrField string
	PtrField *float64
	// Если добавить слайс, то будет ошибка компиляции
	// так как для слайсов не определена операция сравнения
	// SliceField []int
}

func (st1 SomeType) IsEqual(st2 SomeType) bool {
	// Переменные не будут равны, так как прямое сравнение указателей
	// PtrField сравнивает их адреса, а не значения (если оба nil, то равны)
	// Таким образом добавление в структуру ссылочных типов делает
	// невозможным прямое сравнение этих структур
	return st1 == st2
}

func (st1 SomeType) IsDeepEqual(st2 SomeType) bool {
	// DeepEqual сравнивает каждый элемент структуры
	// Однако операция не быстрая из-за рекурсивного обхода
	return reflect.DeepEqual(st1, st2)
}

type ReflectStruct struct {
	FieldName int
}

func main() {
	floatValue1, floatValue2 := 10.0, 10.0
	a := SomeType{IntField: 1, StrField: "str", PtrField: &floatValue1}
	b := SomeType{IntField: 1, StrField: "str", PtrField: &floatValue2}
	fmt.Printf("a и b не равны: %v\n", a.IsEqual(b))
	fmt.Printf("a и b равны: %v\n", a.IsDeepEqual(b))

	// С помощью ValueOf можно определить тип переменной
	// ValueOf возвращает универсальный тип Value
	var varBool *bool
	fmt.Println("ptr:", reflect.ValueOf(varBool).Kind())   // тип объекта
	fmt.Println("*bool:", reflect.ValueOf(varBool).Type()) // оригинальный (не пользовательский) тип
	// Если передан не ссылочный тип, то IsNil() вызовет панику
	fmt.Println("Is nil:", reflect.ValueOf(varBool).IsNil())

	varBoolValue := reflect.ValueOf(varBool)
	// Получить значение, если передан указатель
	fmt.Println("Value of varBoolValue:", varBoolValue.Elem())
	// А здесь будет паника из-за попытки получения значения пустого Value
	//fmt.Println("Value of varBoolValue:", varBoolValue.Elem().Bool())
	trueVal := true
	// Установить значение с помощью рефлексии
	reflect.ValueOf(&varBool).Elem().Set(reflect.ValueOf(&trueVal))
	fmt.Println("With reflect:", reflect.ValueOf(varBool).Elem().Bool())
	fmt.Println("Without reflect:", *varBool)

	// Получаем значение поля по названию (аналог getattr() в Python) или по индексу
	s := ReflectStruct{}
	fieldIndex := 0
	fmt.Println("FieldByName:", reflect.ValueOf(s).FieldByName("FieldName"))
	fmt.Println("FieldByIndex:", reflect.ValueOf(s).Field(fieldIndex))
}
