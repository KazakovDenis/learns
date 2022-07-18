/*
https://go.dev/blog/cover
https://pkg.go.dev/cmd/cover
https://github.com/stretchr/testify
https://laiyuanyuan-sg.medium.com/mock-solutions-for-golang-unit-test-a2b60bd3e157
*/
package foo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

/*
Error, Errorf - Пометить тест как непройденный и записать в error-лог
Fatal, Fatalf - Пометить тест как непройденный и выйти немедленно
Skip, Skipf - Пропустить с сообщением
Log, Logf - Логируют с указанием принадлежности к тесту
*/

func TestFooFunc(t *testing.T) {
	expectedFooResult := "bar"
	if actualFooResult := Foo(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}

// ExampleTestSuite — это набор тестов, который
// создан путём встраивания suite.Suite.
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// TestExampleTestSuite - входная точка сьюта
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// TestAdd - тест с подтестами
func TestAddSubTests(t *testing.T) {
	t.Run("Both zeros", func(t *testing.T) {
		_, err := Add(0, 0)
		require.Error(t, err)
	})

	t.Run("One zero", func(t *testing.T) {
		_, err := Add(0, 1)
		require.Error(t, err)
	})

	t.Run("ZeroDenumenator", func(t *testing.T) {
		res, err := Add(1, 1)
		require.NoError(t, err)
		assert.Equal(t, 2, res)
	})
}

// Table-driven
func TestAddTableDriven(t *testing.T) {
	// аргументы тестируемой функции
	type args struct {
		a int
		b int
	}
	// структура тестовых данных
	tests := []struct {
		name     string
		args     args
		expected int
		expError bool
	}{
		// test cases
		{
			name: "Test Positive",
			args: args{
				a: 1,
				b: 2,
			},
			expected: 3,
			expError: false,
		},
		{
			name: "Test Negative 1",
			args: args{
				a: -1,
				b: 2,
			},
			expected: 0,
			expError: true,
		},
		{
			name: "Test Negative 2",
			args: args{
				a: 1,
				b: -2,
			},
			expected: 0,
			expError: true,
		},

		{
			name: "Test Negative all",
			args: args{
				a: -1,
				b: -2,
			},
			expected: 0,
			expError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.a, tt.args.b)
			if (err != nil) != tt.expError {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.expError)
				return
			}
			if got != tt.expected {
				t.Errorf("Add() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Тестирование интерфейсов
// DBStorage - интерфейс, который надо протестировать
type DBStorage interface {
	UserExists(email string) bool
}

func NewUser(db DBStorage, email string) error {
	if db.UserExists(email) {
		return fmt.Errorf(`user with '%s' email already exists`, email)
	}
	// логика работы
	return nil
}

type DBMock struct {
	emails  map[string]bool
	counter int
}

// Заглушка должна соответствовать интерфейсу
func (db *DBMock) UserExists(email string) bool {
	db.counter++
	return db.emails[email]
}

func (db *DBMock) addUser(email string) {
	db.emails[email] = true
}

func TestNewUser(t *testing.T) {
	errPattern := `user with '%s' email already exists`
	tbl := []struct {
		name     string
		email    string
		preset   bool
		expError bool
	}{
		{`success expected`, `john@test.com`, false, false},
		{`error expected`, `doe@test.com`, true, true},
	}
	for _, item := range tbl {
		t.Run(item.name, func(t *testing.T) {
			dbMock := &DBMock{emails: make(map[string]bool)}
			if item.preset {
				dbMock.addUser(item.email)
			}

			err := NewUser(dbMock, item.email)
			if !item.expError {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, fmt.Sprintf(errPattern, item.email))
			}
		})
	}
}
