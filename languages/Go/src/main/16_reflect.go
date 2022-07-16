// https://pkg.go.dev/reflect
package main

/*
Рефлексия - возможность получить информацию о типе из переменной этого типа.
С помощью пакета reflect можно динамически создавать типы в рантайме.
*/

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
	FieldName string `tag1:"value1" tag2:"value2"`
}

func PrintFieldTags(s *ReflectStruct, fieldIdx int) {
	objType := reflect.ValueOf(s).Type()
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	field := objType.Field(fieldIdx)
	fmt.Println(parseTagString(string(field.Tag)))
}

type TagsInfo map[string][]string

// parseTagString десериализует тег-строку поля структуры.
// Дедупликация имён тегов: первый по порядку (слева направо).
// Ограничения: значение тега не может содержать символы ':' и '"'.
func parseTagString(tagRaw string) (retInfos TagsInfo) {
	retInfos = make(TagsInfo)

	// пример строки: json:"name" pg:"nullable,sortable"
	for _, tag := range strings.Split(tagRaw, " ") {
		if tag = strings.TrimSpace(tag); tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ":")
		if len(tagParts) != 2 {
			continue
		}

		tagName := strings.TrimSpace(tagParts[0])
		if _, found := retInfos[tagName]; found {
			continue
		}

		tagValuesRaw, _ := strconv.Unquote(tagParts[1])
		tagValues := make([]string, 0)
		for _, value := range strings.Split(tagValuesRaw, ",") {
			if value := strings.TrimSpace(value); value != "" {
				tagValues = append(tagValues, value)
			}
		}

		retInfos[tagName] = tagValues
	}
	return
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
	s := ReflectStruct{FieldName: "FieldValue"}
	fieldIndex := 0
	fmt.Println("FieldByName:", reflect.ValueOf(s).FieldByName("FieldName"))
	fmt.Println("FieldByIndex:", reflect.ValueOf(s).Field(fieldIndex))

	// Не все поля структуры можно изменить, проверка осуществляется CanSet()
	// Изменить значения можно только для экспортируемых полей.
	fmt.Println("Can set:", reflect.ValueOf(s).Field(fieldIndex).CanSet())

	// Парсинг тегов
	// Каждому полю структуры можно указать теги с дополнительной информацией
	PrintFieldTags(&s, 0)
}
