package main

import (
	"context"

	"github.com/joho/godotenv"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/KaymeKaydex/link-cut.git/internal/pkg/app"
)

func main() {

	// Установка уровня логирования

	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	// Запуск фонового контекста

	ctx := context.Background()

	// Загружаем переменные окружения

	if err := godotenv.Load(); err != nil {
		jww.ERROR.Println("No .env file loaded")

		return
	}

	// Создание приложения

	a, err := app.NewApp(ctx)
	if err != nil {
		jww.ERROR.Println("Cant create new application : ", err)

		return
	}

	// Запуск приложения
	err = a.Run(ctx)
	if err != nil {
		jww.ERROR.Println("Cant run application : ", err)
		return
	}
}
