package main

import (
	"context"
	"fmt"
	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
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

	// Создаем DSN
	dsn := fmt.Sprintf("host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		"0.0.0.0",
		"5432",
		"user",
		"pass",
		"link-cut")

	// Создание подключения с базой
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
