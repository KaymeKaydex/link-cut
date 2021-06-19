package app

import (
	"context"
	"github.com/KaymeKaydex/link-cut.git/internal/app/repository"
	"github.com/KaymeKaydex/link-cut.git/internal/pkg/server"

	pb "github.com/KaymeKaydex/link-cut.git/pkg/link-cut-api"
	jww "github.com/spf13/jwalterweatherman"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
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

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		jww.FATAL.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterLinkCutServer(s,&server.Server{})

	jww.INFO.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		jww.FATAL.Printf("failed to serve: %v", err)
	}

	return nil
}
