package server

import (
	"context"
	"errors"
	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
	"github.com/KaymeKaydex/link-cut.git/internal/app/repository"
	"github.com/KaymeKaydex/link-cut.git/internal/app/xeger"
	pb "github.com/KaymeKaydex/link-cut.git/pkg/link-cut-api"
	jww "github.com/spf13/jwalterweatherman"
	"net/http"
)

type Server struct {
	r *repository.Repository
	pb.UnimplementedLinkCutServer
}

func NewServer(r * repository.Repository) (*Server,error){
	return &Server{r: r}, nil
}
func (s *Server)Create(ctx context.Context,req *pb.CreateRequest) (*pb.CreateReply,error) {
	jww.DEBUG.Println("New create request :", req.Url)
	longURL:= req.Url
	// Пытаемся постучаться по URL
	if _, err := http.Get(req.GetUrl()); err != nil {
		jww.DEBUG.Println("Invalid url :", longURL)
		return nil, errors.New("invalid url")
	}
	// Смотрим на наличие в базе данных
	if s.r.Exists(ctx, &repository.LongURL{URL: longURL}) {
		return &pb.CreateReply{Url: "URL is already exists"},nil
	}
	// Пытаемся сгенерить рабочую строку
	shortStr := "Number of repetitions exceeded"
	for i:=0;i<100;i++ {
		randStr,err := xeger.GetRandomSequence(5)
		if err != nil{
			jww.INFO.Println("Xeger generate exception")
			continue
		}
		if s.r.Exists(ctx, &repository.ShortURL{URL: "l.ru/"+randStr}){
			jww.DEBUG.Println("New random str exists in DB")
			continue
		}
		err = s.r.CreateContainer(ctx, &model.URLContainer{
			ShortUrl: randStr,
			LongUrl: longURL,
		})

		if err != nil {
			continue
		}
		shortStr = randStr
		break
	}
	return &pb.CreateReply{Url: "l.ru/"+shortStr},nil
}
func (s *Server)Get(ctx context.Context,req *pb.GetRequest) (*pb.GetReply,error) {
	jww.DEBUG.Println("New get request :", req.Url)
	// Пытаемся получить контейнер по сокращенной ссылке
	m, err := s.r.GetContainer(ctx, &repository.ShortURL{
		// Режем ссылку, тк мы не храним в базе данных наше доменное имя
		URL: req.Url[5:],
	})
	if err != nil {
		jww.INFO.Println("Attempt to retrieve a non-existing container")
		return nil, errors.New("attempt to retrieve a non-existing container")
	}
	return &pb.GetReply{
		Url: m.LongUrl,
	}, err
}