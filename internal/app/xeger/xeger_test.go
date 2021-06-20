package xeger

import (
	"testing"
)

func TestGetRandomSequence(t *testing.T) {
	testTable := []struct {
		length    int
		errExists bool
	}{
		{
			length:    5,
			errExists: false,
		},
		{
			length:    -1,
			errExists: true,
		},
		{
			length:    0,
			errExists: true,
		},
		{
			length:    1,
			errExists: false,
		},
	}
	for _, testCase := range testTable {

		r, err := GetRandomSequence(testCase.length)

		// Проверка на наличие ошибок для граничных значений
		if (err == nil) == testCase.errExists {
			t.Errorf("No error for invalid input value %d", testCase.length)
		}

		// Проверка на соответсвие длинны запроса и результата
		if testCase.length > 0 && testCase.length != len(r) {
			t.Errorf("Invalid length of returned value, expected : %d, result: %d", testCase.length, len(r))
		}
	}
}
