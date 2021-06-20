package main

import (
	"context"
	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
	"github.com/KaymeKaydex/link-cut.git/internal/app/repository"
	"github.com/joho/godotenv"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Создает миграции в базу данных

func main() {

	// Устанавливаем уровень логирования
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)
	jww.INFO.Println("Migration started")

	// Запуск фонового контекста
	ctx := context.Background()

	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		jww.ERROR.Println("No .env file loaded")
		return
	}

	// Создание подключения с базой
	db, err := gorm.Open(postgres.Open(repository.GetDSN()), &gorm.Config{})
	if err != nil {
		jww.INFO.Println("Cant open postgers connection", err)
		return
	}

	// Создание миграции
	err = db.WithContext(ctx).AutoMigrate(&model.URLContainer{})
	if err != nil {
		jww.INFO.Println("Migration error :", err)
		return
	}

	jww.INFO.Println("Successfully migrate")
}
