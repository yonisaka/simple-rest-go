package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_http_article "simple-rest-go/internal/app/article/delivery/http"
	_repo_article "simple-rest-go/internal/app/article/repository"
	_ucase_article "simple-rest-go/internal/app/article/usecase"
	"simple-rest-go/internal/app/mysql"
)

func main() {
	r := gin.Default()
	// set default enviroment
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_NAME", "simple-rest-go")

	viper.BindEnv("DB_HOST", "DB_HOST")
	viper.BindEnv("DB_PORT", "DB_PORT")
	viper.BindEnv("DB_USER", "DB_USER")
	viper.BindEnv("DB_PASS", "DB_PASS")
	viper.BindEnv("DB_NAME", "DB_NAME")

	db := mysql.Connection()
	timeoutContext := time.Duration(5 * time.Second)
	repo_article := _repo_article.NewArticleRepository(db)
	ucase_article := _ucase_article.NewArticleUsecase(repo_article, timeoutContext)
	_http_article.NewArticleHandler(r, ucase_article)
	r.Run("localhost:8080")
}
