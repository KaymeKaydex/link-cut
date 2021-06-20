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

const (
	DomainName = "l.ru"
	RetryCount = 100
)

type Server struct {
	r *repository.Repository
	pb.UnimplementedLinkCutServer
}

func NewServer(r *repository.Repository) (*Server, error) {
	return &Server{r: r}, nil
}
func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	jww.INFO.Println("New create request :", req.Url)

	// Записываем входящий длинный URL
	longURL := req.Url

	// Пытаемся постучаться по URL чтобы проверить его релевантность
	if _, err := http.Get(req.GetUrl()); err != nil {
		return nil, errors.New("invalid url")
	}

	// Смотрим на наличие в базе данных
	if s.r.Exists(ctx, &repository.LongURL{URL: longURL}) {
		return nil, errors.New("URL is already exists")
	}

	// Пытаемся сгенерить последовательность
	shortStr := ""
	for i := 0; i < RetryCount; i++ {
		randStr, err := xeger.GetRandomSequence(5)
		if err != nil {
			jww.INFO.Println("Xeger generate exception")
			continue
		}
		if s.r.Exists(ctx, &repository.ShortURL{URL: DomainName + "/" + randStr}) {
			jww.INFO.Println("New random str exists in DB")
			continue
		}

		// Если мы сгенерировали рабочую последовательноть записываем в базу

		err = s.r.CreateContainer(ctx, &model.URLContainer{
			ShortUrl: randStr,
			LongUrl:  longURL,
		})

		if err != nil {
			jww.INFO.Println("Create container exception : ", err)
			continue
		}

		// Если генерация прошла успешно записываем в возвращаемую строку и завершаем цикл

		shortStr = randStr
		break
	}
	return &pb.CreateReply{Url: DomainName + "/" + shortStr}, nil
}
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	jww.DEBUG.Println("New get request :", req.Url)

	// Пытаемся получить контейнер по сокращенной ссылке

	m, err := s.r.GetContainer(ctx, &repository.ShortURL{

		// Режем ссылку, тк мы не храним в базе данных наше доменное имя ( отрезаем "l.ru/" )

		URL: req.Url[len(DomainName)+1:],
	})
	if err != nil {
		jww.INFO.Println("Attempt to retrieve a non-existing container")
		return nil, errors.New("attempt to retrieve a non-existing container")
	}
	return &pb.GetReply{
		Url: m.LongUrl,
	}, err
}
