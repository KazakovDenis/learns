package main

var condition bool
var x, y int
var caseVar = 5

func main() {
	if !condition {
		println("Flag is false")
	}
	if x == 1 && y == 2 {
	}

	// a и b объявлены в 01_data_types.go
	a := 0.10000001 // float64
	// ветвление с инициализацией
	if b := float32(a); b > float32(0.1) {
		println("Var a is GT float32(0.1)")
	}

	switch caseVar {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3, 4:
		println("3 or 4")
	default:
		println("Default case")
	}

	// Переменную условия можно не задавать
	switch {
	case caseVar > 0:
		println("No case var")
	}

	// Или создать переменную локальной области
	switch anotherVar := caseVar; {
	case anotherVar > 0:
		println("Another case var")
	}

	caseVar = -100
	switch {
	case caseVar > 0:
		if caseVar%2 == 0 {
			// Выйти из switch-case
			break
		}
		println("Odd positive value received")
	case caseVar < 0:
		println("Negative value received")
		// Продолжить выполнение следующих switch-case
		// Но условие следующего case игнорирутеся
		fallthrough
	case caseVar == -10:
		println("Case ignored")
	default:
		println("Default value handling")
	}
}
