package server

import (
	"context"
	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
	"github.com/KaymeKaydex/link-cut.git/internal/app/repository"
	"github.com/KaymeKaydex/link-cut.git/internal/app/xeger"
	pb "github.com/KaymeKaydex/link-cut.git/pkg/link-cut-api"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

type Server struct {
	pb.UnimplementedLinkCutServer
}

func (s *Server)Create(ctx context.Context,req *pb.CreateRequest) (*pb.CreateReply,error) {

	db, err := gorm.Open(postgres.Open(repository.GetDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}) 	// Написатьь отдельную функцию для формирования строки dsn
	if err!=nil {
		jww.ERROR.Println("Cant open postgers connection",err)
		return nil, err
	}
	longURL := req.GetUrl()
	jww.INFO.Println("New Create req :", longURL)
	// Пытаемся постучаться по URL
	if _, err := http.Get(req.GetUrl()); err!=nil {
		jww.DEBUG.Println("Invalid url :", longURL)
		return &pb.CreateReply{Url: "Invalid url :"},nil
	}

	shortStr := ""
	for ;; {
		randStr,err := xeger.GetRandomSequence(5)
		if err != nil{ continue }
		err = db.WithContext(ctx).Create(&model.URLContainer{
			ShortUrl: randStr,
			LongUrl: longURL,
		}).Error
		if err != nil {
			continue
		}
		shortStr = randStr
		break
	}


	return &pb.CreateReply{Url: "l.ru/"+shortStr},nil
}
func (s *Server)Get(ctx context.Context,req *pb.GetRequest) (*pb.GetReply,error) {
	db, err := gorm.Open(postgres.Open(repository.GetDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}) 	// Написатьь отдельную функцию для формирования строки dsn
	if err!=nil {
		jww.ERROR.Println("Cant open postgers connection",err)
		return nil, err
	}
	if len(req.Url)!=10 {
		return &pb.GetReply{Url: "Bad url"}, nil
	}
	m := &model.URLContainer{ShortUrl: req.Url[5:]}
	var foundContainer model.URLContainer
	err = db.WithContext(ctx).Where(m).First(&foundContainer).Error
	if err!= nil {
		return &pb.GetReply{Url: "Not found"}, nil
	}

	return &pb.GetReply{Url: foundContainer.LongUrl}, nil
}