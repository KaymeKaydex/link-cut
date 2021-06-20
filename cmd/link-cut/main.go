package main

import (
	"context"
	"github.com/KaymeKaydex/link-cut.git/internal/pkg/app"
	"github.com/joho/godotenv"
	jww "github.com/spf13/jwalterweatherman"
)

func main() {

	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		jww.ERROR.Println("No .env file loaded")
		return
	}

	a, err := app.NewApp(ctx)
	if err != nil {
		jww.ERROR.Println("Cant create new application : ", err)
		return
	}

	err = a.Run(ctx)
	if err != nil {
		jww.ERROR.Println("Cant run application : ", err)
		return
	}
}
