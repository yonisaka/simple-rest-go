package http

import (
	"log"
	"net/http"
	"simple-rest-go/internal/app/domain"
	"simple-rest-go/internal/app/utils"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	ArticleUcase domain.ArticleUsecase
}

func NewArticleHandler(r *gin.Engine, ucase domain.ArticleUsecase) {
	handler := ArticleHandler{
		ArticleUcase: ucase,
	}
	article := r.Group("/article")

	article.GET("/", handler.getArticles)
}

func (a *ArticleHandler) getArticles(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)

	articleList, err := a.ArticleUcase.Fetch(c, pagination)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": articleList,
	})
}