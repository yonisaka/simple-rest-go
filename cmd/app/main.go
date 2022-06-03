package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_http_article "simple-rest-go/internal/app/article/delivery/http"
	_repo_article "simple-rest-go/internal/app/article/repository"
	_ucase_article "simple-rest-go/internal/app/article/usecase"
	"simple-rest-go/internal/app/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	r := gin.Default()

	db := mysql.Connection()
	timeoutContext := time.Duration(5 * time.Second)
	repo_article := _repo_article.NewArticleRepository(db)
	ucase_article := _ucase_article.NewArticleUsecase(repo_article, timeoutContext)
	_http_article.NewArticleHandler(r, ucase_article)
	r.Run("localhost:8080")
}
