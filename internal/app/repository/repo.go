package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/KaymeKaydex/link-cut.git/internal/app/model"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// CreateContainer на основе объекта модели создает запись в базе данных
func (r *Repository) CreateContainer(ctx context.Context, c *model.URLContainer) error {
	return r.db.WithContext(ctx).Create(c).Error
}

// Exists на основе интерфейса определяте существование строки в таблице
func (r *Repository) Exists(ctx context.Context, url IURL) bool {
	return url.Exists(ctx, r)
}

// На основе интерфейса возвращает строку из таблицы

func (r *Repository) GetContainer(ctx context.Context, url IURL) (*model.URLContainer, error) {
	return url.GetContainer(ctx, r)
}

type IURL interface {
	Exists(ctx context.Context, r *Repository) bool
	GetContainer(ctx context.Context, r *Repository) (*model.URLContainer, error)
}

// ShortURL is container for SHORT URL
type ShortURL struct {
	URL string
}

func (url *ShortURL) Exists(ctx context.Context, r *Repository) bool {
	var foundContainer model.URLContainer

	err := r.db.WithContext(ctx).Where(&model.URLContainer{ShortUrl: url.URL}).First(&foundContainer).Error
	if err != nil {
		return false
	}

	return true
}

func (url *ShortURL) GetContainer(ctx context.Context, r *Repository) (*model.URLContainer, error) {
	var foundContainer model.URLContainer

	err := r.db.WithContext(ctx).Where(&model.URLContainer{ShortUrl: url.URL}).First(&foundContainer).Error
	if err != nil {
		return nil, err
	}

	return &foundContainer, nil
}

// For LONG URL

type LongURL struct {
	URL string
}

func (url *LongURL) Exists(ctx context.Context, r *Repository) bool {
	var foundContainer model.URLContainer

	err := r.db.WithContext(ctx).Where(&model.URLContainer{LongUrl: url.URL}).First(&foundContainer).Error
	if err != nil {
		return false
	}

	return true
}
func (url *LongURL) GetContainer(ctx context.Context, r *Repository) (*model.URLContainer, error) {
	var foundContainer model.URLContainer

	err := r.db.WithContext(ctx).Where(&model.URLContainer{LongUrl: url.URL}).First(&foundContainer).Error
	if err != nil {
		return nil, err
	}

	return &foundContainer, nil
}
