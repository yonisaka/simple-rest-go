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

func (a *articleUsecase) Store(c context.Context, article *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err = a.articleRepo.Store(ctx, article)
	return
}

func (a *articleUsecase) Update(ctx context.Context, article *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.articleRepo.Update(ctx, article)
	return
}

func (a *articleUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.articleRepo.Delete(ctx, id)
	return
}
