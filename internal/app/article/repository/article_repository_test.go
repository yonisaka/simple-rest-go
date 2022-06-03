package repository_test

import (
	"context"
	"fmt"
	"simple-rest-go/internal/app/domain"
	"testing"

	articleRepo "simple-rest-go/internal/app/article/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "localhost", "3306", "simple-rest-go")
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = gdb
}

func TestFetch(t *testing.T) {
	a := articleRepo.NewArticleRepository(db)
	p := domain.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "created_at",
	}
	list, err := a.Fetch(context.TODO(), p)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	a := articleRepo.NewArticleRepository(db)
	id := int64(1)

	res, err := a.GetByID(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCreate(t *testing.T) {
	a := articleRepo.NewArticleRepository(db)
	article := domain.Article{
		Title:    "Test",
		Content:  "Test",
		AuthorID: 1,
	}

	err := a.Store(context.TODO(), &article)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	a := articleRepo.NewArticleRepository(db)
	article := domain.Article{
		ID:       14,
		Title:    "Test",
		Content:  "Test",
		AuthorID: 1,
	}

	err := a.Update(context.TODO(), &article)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	a := articleRepo.NewArticleRepository(db)
	id := int64(14)

	err := a.Delete(context.TODO(), id)
	assert.NoError(t, err)
}
