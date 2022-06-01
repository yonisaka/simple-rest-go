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
