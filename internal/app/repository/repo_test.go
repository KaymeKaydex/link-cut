package repository

import (
	"context"
	"fmt"
	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"testing"
)

var dsn string // Хранит DSN для подключения к базе данных

func init() {
	dsn = fmt.Sprintf("host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		"0.0.0.0",
		"5432",
		"user",
		"pass",
		"link-cut")
}

func TestLongURL_Exists(t *testing.T) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Написатьь отдельную функцию для формирования строки dsn
	if err != nil {
		jww.ERROR.Fatalln("Cant open posters connection", err)
	}
	r := New(db)

	example := &model.URLContainer{
		LongUrl:  "https://yandex.ru/",
		ShortUrl: "yyyyy",
	}
	err = r.CreateContainer(ctx, example)
	if err != nil {
		jww.ERROR.Fatalln("Cant create container for example ", err)
	}
	result := r.Exists(ctx, &LongURL{URL: example.LongUrl})

	if result != true {
		t.Error("Bad Long URL exists")
	}
	if err = db.WithContext(ctx).Delete(example).Error; err != nil {
		jww.FATAL.Fatalln("Cant delete example", err)
	}
}

func TestShortURL_Exists(t *testing.T) {

	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Написатьь отдельную функцию для формирования строки dsn
	if err != nil {
		jww.ERROR.Fatalln("Cant open posters connection", err)
	}
	r := New(db)

	example := &model.URLContainer{
		LongUrl:  "https://google.ru/",
		ShortUrl: "ggggg",
	}
	err = r.CreateContainer(ctx, example)
	if err != nil {
		jww.ERROR.Fatalln("Cant create container for example ", err)
	}
	result := r.Exists(ctx, &ShortURL{URL: example.ShortUrl})

	if result != true {
		t.Error("Bad Short URL exists")
	}
	if err = db.WithContext(ctx).Delete(example).Error; err != nil {
		jww.FATAL.Fatalln("Cant delete example", err)
	}
}

func TestLongURL_GetContainer(t *testing.T) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Написатьь отдельную функцию для формирования строки dsn
	if err != nil {
		jww.ERROR.Fatalln("Cant open posters connection", err)
	}
	r := New(db)

	example := &model.URLContainer{
		LongUrl:  "https://mail.ru/",
		ShortUrl: "mailr",
	}
	err = r.CreateContainer(ctx, example)

	if err != nil {
		jww.ERROR.Fatalln("Cant create container for example ", err)
	}
	result, err := r.GetContainer(ctx, &LongURL{URL: example.LongUrl})

	if result.LongUrl != example.LongUrl && result.ShortUrl != example.ShortUrl {
		t.Errorf("Incorrect GetContainer by long url. Expected : %s, result: %s ", example, result)
	}
	if err = db.WithContext(ctx).Delete(example).Error; err != nil {
		jww.FATAL.Fatalln("Cant delete example", err)
	}
}
func TestShortURL_GetContainer(t *testing.T) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Написатьь отдельную функцию для формирования строки dsn
	if err != nil {
		jww.ERROR.Fatalln("Cant open posters connection", err)
	}
	r := New(db)

	example := &model.URLContainer{
		LongUrl:  "https://ozon.ru/",
		ShortUrl: "ooooo",
	}
	err = r.CreateContainer(ctx, example)

	if err != nil {
		jww.ERROR.Fatalln("Cant create container for example ", err)
	}
	result, err := r.GetContainer(ctx, &ShortURL{URL: example.ShortUrl})

	if result.LongUrl != example.LongUrl && result.ShortUrl != example.ShortUrl {
		t.Errorf("Incorrect GetContainer by long url. Expected : %s, result: %s ", example, result)
	}
	if err = db.WithContext(ctx).Delete(example).Error; err != nil {
		jww.FATAL.Fatalln("Cant delete example", err)
	}
}
