package main

import (
	"fmt"
)

func checkSlice(s []string) {
	s[0] = "a"
	b := append(s, "b")
	fmt.Println("Slice in func:", b)
}

func main() {
	// ARRAYS
	// Массив на 7 целых положительных чисел (по умолчанию нулей)
	// Количество элементов в массиве — это часть типа,
	// то есть массивы [3]int и [5]int относятся к разным типам
	var emptyIntArray [7]uint8
	fmt.Println(emptyIntArray[0])

	// С инициализацией
	nonEmptyIntArray := [7]int{-3, 5, 7} // [-3 5 7 0 0 0 0]
	fmt.Println(nonEmptyIntArray[1])

	// Размер определён по кол-ву инициализирующих элементов
	rgbColor := [...]uint8{255, 255, 128} // len = 3
	fmt.Println(len(rgbColor))

	// Инициализация элементов по индексу (индекс:значение)
	thisWeekTemp := [7]int{6: 11, 2: 3} // [0 0 3 0 0 0 0 11]
	fmt.Println(thisWeekTemp[2])

	// Многомерный массив
	// var rgbImage[1080][1920][3]uint8

	// Обход массива
	var weekTemp = [4]int{5, 4, 6}
	// Обращаемся по указателю, чтобы избежать копирования
	for idx, temp := range &weekTemp {
		fmt.Println(idx, temp)
	}

	// SLICES
	// Слайс - коллекция переменной длины, который представляет собой:
	// Массив + Заголовок с указателем на него, длиной + ёмкостью
	// var Array [2]int
	var Slice1 []int // равен nil
	fmt.Println("Возможная длина", cap(Slice1), "Занятая длина", len(Slice1))

	// Для создания слайса используется встроенная функция make()
	// Slice := make([]TypeOfelement, LenOfslice, CapOfSlice)
	// mySlice := make([]int)   // слайс [], базовый массив []
	// mySlice := make([]int,5) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	Slice2 := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(Slice2[0])

	// С инициализацией, cap == len
	Slice3 := []uint8{1, 2, 3}
	fmt.Println("Init", len(Slice3), cap(Slice3))

	// Слайс из массива или другого слайса
	Slice4 := weekTemp[0:]
	fmt.Println("Slice from slice", Slice4[0])

	// Базовый массив тоже изменяется
	// Присвоить значение по индексу >= cap нельзя
	Slice4[0] = 11
	fmt.Println("Base array changed", Slice4[0], "==", weekTemp[0])

	// append создаёт новый слайс с новыми элементами, а не изменяет существующий
	// Механика как в Python - если len >= cap, то слайс создаётся с новым
	// базовым массивом вдвое больше текущего
	// Из-за неявности формирования, рекомендуют append использовать для
	// присвоения слайса самому себе: s = append(s, elem)
	newSliceFrom4 := append(Slice4, 35, 36)
	fmt.Println("Append:", len(Slice4), len(newSliceFrom4))

	// Соединение двух слайсов
	mergedSlice := append(Slice1, newSliceFrom4...)
	fmt.Println("Merged slice", mergedSlice)

	// append в слайсе добавит только в локальный слайс стек функции
	// присвоение по индексу изменит и исходный
	funcSlice := make([]string, 3, 5)
	checkSlice(funcSlice)
	fmt.Println("Slice after func", funcSlice)
}
