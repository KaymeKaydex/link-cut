package xeger

import (
	"errors"
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/takahiromiyamoto/go-xeger"
)

// GetRandomSequence генерация случайной последовательности из n символов, с использованием regexp в библиотеке Xeger
func GetRandomSequence(length int) (string, error) {
	if length < 1 {
		jww.INFO.Println("Invalid length : ", length)

		return "", errors.New("Invalid length")
	}

	x, err := xeger.NewXeger(fmt.Sprintf("[a-zA-Z0-9\\_]{%d}", length))
	if err != nil {
		jww.INFO.Println("Cant generate sequence", err)
		return "", err
	}

	return x.Generate(), nil
}
