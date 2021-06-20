package xeger

import (
	"testing"
)

func TestGetRandomSequence(t *testing.T) {
	testTable := []struct {
		lenght    int
		errExists bool
	}{
		{
			lenght:    5,
			errExists: false,
		},
		{
			lenght:    -1,
			errExists: true,
		},
		{
			lenght:    0,
			errExists: true,
		},
		{
			lenght:    1,
			errExists: false,
		},
	}
	for _, testCase := range testTable {

		r, err := GetRandomSequence(testCase.lenght)

		// Проверка на наличие ошибок для граничных значений
		if (err == nil) == testCase.errExists {
			t.Errorf("No error for invalid input value %d", testCase.lenght)
		}

		// Проверка на соответсвие длинны запроса и результата
		if testCase.lenght > 0 && testCase.lenght != len(r) {
			t.Errorf("Invalid lenght of returned value, expected : %d, result: %d", testCase.lenght, len(r))
		}
	}
}
