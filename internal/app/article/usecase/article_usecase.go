package usecase

import (
	"context"
	"simple-rest-go/internal/app/domain"
	"time"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(article domain.ArticleRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    article,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) Fetch(c context.Context, pagination domain.Pagination) (res []domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.Fetch(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return
}

func (a *articleUsecase) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return res, err
	}

	return
}

func (a *articleUsecase) Store(c context.Context, m *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err = a.articleRepo.Store(ctx, m)
	return
}
