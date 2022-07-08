package main

import "fmt"

func fizzBuzz(count int) {
	i := 1
	for i <= count {
		switch {
		case i%3 == 0 && i%5 == 0:
			println("FizzBuzz")
			break
		case i%3 == 0:
			println("Fizz")
		case i%5 == 0:
			println("Buzz")
		default:
			println(i)
		}
		i++
	}
}

func main() {
	// Как и везде: break, continue

	// Бесконечный цикл
	// for {}

	// Трёхкомпонентный цикл
	for i := 1; i < 2; i++ {
		// наращиваем переменную
		println("For", i)
	}

	// Бесконечные трёхкомпонетные циклы
	// for ;; {}
	// for ; true; {}

	// Аналог while
	i := true
	for i != false {
		println("While", i)
		i = false
	}

	// Итерация по массиву (цикл range)
	array := [3]int{1, 2, 3}
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}

	// Метки (labels) и goto
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			// Перейти к следующей итерации внешнего цикла
			// continue outerLoopLabel
			// Выйти из внешнего цикла
			break outerLoopLabel
		}
		goto outerLoopLabel
	}
	fmt.Println("End")

	fizzBuzz(20)
}
