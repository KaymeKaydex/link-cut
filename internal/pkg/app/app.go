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
	s *server.Server
}

func NewApp(_ context.Context) (*App, error) {
	db, err := gorm.Open(postgres.Open(repository.GetDSN()), &gorm.Config{}) // Написатьь отдельную функцию для формирования строки dsn
	if err != nil {
		jww.ERROR.Println("Cant open postgers connection", err)
		return nil, err
	}
	r := repository.New(db)
	serv, err := server.NewServer(r)
	if err != nil {
		jww.FATAL.Println("Cant create server")
		return nil, err
	}
	return &App{
		s: serv,
	}, nil

}

func (a *App) Run(_ context.Context) error {

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		jww.FATAL.Printf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()

	pb.RegisterLinkCutServer(s, a.s)

	jww.INFO.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		jww.FATAL.Printf("failed to serve: %v", err)
	}

	return nil
}
