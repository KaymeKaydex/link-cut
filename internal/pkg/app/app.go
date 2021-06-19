package app

import (
	"context"
	"github.com/KaymeKaydex/link-cut.git/internal/app/repository"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	r *repository.Repository
}

func NewApp(ctx context.Context) (*App,error) {
	db, err := gorm.Open(postgres.Open(repository.GetDSN()), &gorm.Config{}) 	// Написатьь отдельную функцию для формирования строки dsn
	if err!=nil {
		jww.ERROR.Println("Cant open postgers connection",err)
		return nil, err
	}

	return &App{
		r: repository.New(db),
	},nil

}

func (a *App) Run(ctx context.Context) error {


	return nil
}
