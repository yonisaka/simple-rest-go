package repository

import (
	"context"
	"simple-rest-go/internal/app/domain"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.ArticleRepository {
	return &articleRepository{db}
}

func (a *articleRepository) Fetch(ctx context.Context, pagination domain.Pagination) (res []domain.Article, err error) {
	var articles []domain.Article
	offset := (pagination.Page - 1) * pagination.Limit
	query := a.db.Limit(pagination.Limit).
		Offset(offset).
		Order(pagination.Sort)

	result := query.Model(&domain.Article{}).
		// Preload("Author").
		Joins("Author").
		Find(&articles)

	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return articles, nil
}

func (a *articleRepository) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	result := a.db.Joins("Author").First(&res, id)
	if result.Error != nil {
		msg := result.Error
		return res, msg
	}
	return res, nil
}

func (a *articleRepository) Store(ctx context.Context, article *domain.Article) (err error) {
	result := a.db.Create(&article)
	if result.Error != nil {
		msg := result.Error
		return msg
	}
	return
}

func (a *articleRepository) Update(ctx context.Context, article *domain.Article) (err error) {
	result := a.db.Model(&article).Updates(article)
	if result.Error != nil {
		msg := result.Error
		return msg
	}
	return
}

func (a *articleRepository) Delete(ctx context.Context, id int64) (err error) {
	result := a.db.Delete(&domain.Article{}, id)
	if result.Error != nil {
		msg := result.Error
		return msg
	}
	return
}
