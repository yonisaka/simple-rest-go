package usecase_test

import (
	"context"
	"errors"
	"simple-rest-go/internal/app/article/usecase"
	"simple-rest-go/internal/app/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	InternalError = errors.New("internal error")
)

type fields struct {
	article domain.ArticleUsecase
	timeout time.Duration
}

func sut(f fields) domain.ArticleUsecase {
	return usecase.NewArticleUsecase(f.article, f.timeout)
}

func TestFetch(t *testing.T) {
	type args struct {
		ctx        context.Context
		pagination domain.Pagination
	}

	type test struct {
		fields  fields
		args    args
		want    []domain.Article
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling Article Fetch succeed, Should return no error": func(t *testing.T) test {
			ctx := context.Background()
			pagination := domain.Pagination{
				Page:  1,
				Limit: 2,
				Sort:  "created_at",
			}

			args := args{
				ctx:        ctx,
				pagination: pagination,
			}

			return test{
				fields: fields{
					article: &domain.ArticleUsecaseMock{
						FetchFunc: func(ctx context.Context, pagination domain.Pagination) ([]domain.Article, error) {
							return []domain.Article{
								{
									ID:        1,
									Title:     "Title",
									Content:   "Content",
									CreatedAt: time.Now(),
									UpdatedAt: time.Now(),
								},
							}, nil
						},
					},
				},
				args: args,
				want: []domain.Article{
					{
						ID:        1,
						Title:     "Title",
						Content:   "Content",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				wantErr: nil,
			}
		},
		"Given valid request parameter, When calling Ipify Service failed, Should return error": func(t *testing.T) test {
			ctx := context.Background()
			pagination := domain.Pagination{
				Page:  1,
				Limit: 2,
				Sort:  "created_at",
			}

			args := args{
				ctx:        ctx,
				pagination: pagination,
			}

			return test{
				fields: fields{
					article: &domain.ArticleUsecaseMock{
						FetchFunc: func(ctx context.Context, pagination domain.Pagination) ([]domain.Article, error) {
							return []domain.Article{}, InternalError
						},
					},
				},
				args:    args,
				want:    []domain.Article(nil),
				wantErr: InternalError,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			got, err := sut.Fetch(tt.args.ctx, tt.args.pagination)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
