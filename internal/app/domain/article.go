package domain

import (
	"context"
	"time"
)

//go:generate rm -f ./article_usecase_mock.go
//go:generate moq -out ./article_usecase_mock.go . ArticleUsecase:ArticleUsecaseMock

type Article struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(100);not null" validate:"required"`
	Content   string    `json:"content" gorm:"type:text;not null" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	Author    Author    `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type ArticleUsecase interface {
	Fetch(ctx context.Context, pagination Pagination) ([]Article, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Store(context.Context, *Article) error
	Update(ctx context.Context, ar *Article) error
	Delete(ctx context.Context, id int64) error
}

type ArticleRepository interface {
	Fetch(ctx context.Context, pagination Pagination) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Store(ctx context.Context, a *Article) error
	Update(ctx context.Context, ar *Article) error
	Delete(ctx context.Context, id int64) error
}
