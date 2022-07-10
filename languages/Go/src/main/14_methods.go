package main

import "fmt"

type MyType int

// Метод принимает тип, к которому относится, в качестве
// аргумента (получателя) перед объявлением названия метода
// Как видно, методы могут быть не только у структур
func (m MyType) String() string {
	return fmt.Sprintf("MyType = %d", m)
}

// Метод может принимать получателя как по значению, так и по указателю
// Вызов метода с получателем по значению для указателя на объект будет эквивалентен вызову (*b).Method().
// Вызов метода с получателем по указателю для значения объекта будет эквивалентен вызову (&b).Method().
// Как правило, придерживаются одного из этих вариантов.

// CircularBuffer реализует структуру данных «кольцевой буфер» для значений float64.
type CircularBuffer struct {
	values  []float64 // Текущие значения буфера
	headIdx int       // Индекс головы (первый непустой элемент)
	tailIdx int       // Индекс хвоста (первый пустой элемент)
}

// GetCurrentSize возвращает текущую длину буфера.
// Value receiver - получает значение (копию структуры)
// Но при изменении ссылочного типа (values)
// также будет изменяться и оригинальный
func (b CircularBuffer) GetCurrentSize() int {
	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.values) - b.headIdx
	}
	return b.tailIdx - b.headIdx
}

// AddValue добавляет новое значение в буфер.
// Pointer receiver - получает указатель на структуру
func (b *CircularBuffer) AddValue(v float64) {
	b.values[b.tailIdx] = v
	b.tailIdx = (b.tailIdx + 1) % cap(b.values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.values)
	}
}

// NewCircularBuffer — конструктор типа CircularBuffer.
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{values: make([]float64, size+1)}
}

// LogStruct
// Т.к. функции в Go первого порядка, они могут быть полями структур
type LogStruct struct {
	// Однако такие функции не имеют доступа к владельцу,
	// если он не был передан явно, а также может быть переопределена
	Log func(s string)
}

// Наследования нет, но структуры могут использовать методы вложенных в них

// Person — структура, описывающая человека.
type Person struct {
	Name string
	Year int
}

// String возвращает информацию о человеке.
func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке.
func (p Person) Print() {
	fmt.Println(p.String())
}

// NewPerson возвращает новую структуру Person.
func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

// Student описывает студента с использованием вложенной структуры Person. То есть структура Student описывает.
// У структуры Student нет метода Print, но она может его использовать через Person
// Грубо говоря, Student "наследует" все поля и методы Person
type Student struct {
	// Если указать Person Person, структура уже не будет вложенной
	// Тип может встраиваться и по указателю с соответствующими последствиями: *Person
	// Нельзя встраивать указатели на указатели: **Person
	Person
	Group string
}

// String возвращает информацию о студенте.
func (s Student) String() string {
	// у Person будет также вызван метод String()
	// Работает похожим образом MRO: сначала Go ищет метод
	// у самой структуры, затем у вложенных в неё структур
	// Если >2 вложенных структур имеют одинаковые методы, происходит конфликт
	// Ошибка будет только в случае, если этот метод используется в коде.
	// В случае конфликтов следует явно указывать у кого вызывать метод,
	// например: student.Person.String()
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func main() {
	var m MyType = 5
	var s = m.String()
	fmt.Println("String from method:", s)

	buf := NewCircularBuffer(4)
	buf.AddValue(1.23)
	fmt.Println("Buffer size:", buf.GetCurrentSize())

	var logStruct = LogStruct{
		Log: func(s string) { fmt.Println(s) },
	}
	logStruct.Log("Logged")

	student := Student{Person: NewPerson("Ivan", 2022), Group: "This"}
	// Мы можем переопределить у Student метод String()
	fmt.Println("student", student.String())
}
