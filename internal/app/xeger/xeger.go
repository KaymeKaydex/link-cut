package xeger

import (
	"errors"
	"fmt"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/takahiromiyamoto/go-xeger"
)

// Генерация случайной последовтельности из n символов, с использованием regexp

func GetRandomSequence(lenght int) (string, error){
	if lenght <1 {
		jww.INFO.Println("Invalid lenght : ",lenght)
		return "", errors.New("Invalid lenght")
	}
	x, err := xeger.NewXeger(fmt.Sprintf("[a-zA-Z0-9\\_]{%d}",lenght))
	if err != nil {
		jww.INFO.Println("Cant generate sequence", err)
		return "",err
	}
	return x.Generate(),nil
}