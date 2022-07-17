package main

import (
	"errors"
	"fmt"
	"os"

	extErrors "github.com/pkg/errors"
)

/*
Тип error - встроенный интерфейсный тип с методом Error(),
который должен возвращать текст ошибки.

Если ошибка, вернувшаяся из функции == nil, значит функция
отработала корректно.

Ошибки можно сравнивать между собой, но нельзя сравнивать
ошибки, сформированные динамически, ни с чем, кроме nil.

Хорошей практикой является начинать текст ошибки с названия
пакета, где она объявлена, так будет проще её найти
*/

var someError = errors.New("package_name: some bad error %s")

func wrapError() error {
	_, err := os.ReadFile("NoSuchFile.txt")
	// Спецификатор %w оборачивает исходную ошибку
	return fmt.Errorf(`не удалось прочитать файл (%w)`, err)
}

/* Panic
В аварийной ситуации программа останавливает работу, вызываются defer
и выводится сообщение об ошибке со стеком вызовов. В функцию panic()
можно передать значение любого типа.

Не рекомендуется использовать функцию panic для обработки всех ошибок,
так как этот механизм требует ресурсов для раскручивания стека.
Если паники всё же вызываются, то они должны быть тщательно задокументированы.

Если произошла паника, то восстановить состояние программы можно
вызовом функции recover(), которая остановит панику и вернёт аргумент,
переданный в неё либо nil, т.о. передавать nil в панику не следует.
Применять recover() следует лишь в оправданных ситуациях.
*/
func somePanic() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(`Возникла паника: `, p)
		}
	}()
	panic(`aварийная ситуация`)
}

func main() {
	// Динамически созданная ошибка
	newError := fmt.Errorf("package_name: new error %s", "happened")
	fmt.Println(someError)
	fmt.Println(newError)

	wrappedError := wrapError()
	fmt.Println(wrappedError)

	originalError := errors.Unwrap(wrappedError)
	fmt.Println(originalError)

	wrapAgain := extErrors.Wrap(wrappedError, "wrapped again")
	fmt.Println(wrapAgain)
	// Является ли wrappedError обёрткой для originalError?
	fmt.Println("Сравнение ошибок:", errors.Is(wrappedError, originalError))
	fmt.Println("Сравнение ошибок:", errors.Is(wrappedError, wrapAgain))

	// As находит первую в цепочке ошибку err, устанавливает
	// тип, равным этому значению ошибки, и возвращает true
	fmt.Println("Приведение ошибок:", errors.As(wrappedError, &wrapAgain))

	somePanic()
	fmt.Println("Программа завершилась корректно")
}
