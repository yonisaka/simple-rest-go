package http

import (
	"net/http"
	"simple-rest-go/internal/app/domain"
	"simple-rest-go/internal/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
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
	article.GET("/:id", handler.getArticleByID)
	article.POST("/", handler.store)
}

func (a *ArticleHandler) getArticles(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)

	articleList, err := a.ArticleUcase.Fetch(c, pagination)

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

func (a *ArticleHandler) getArticleByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": utils.ErrBadParamInput.Error(),
		})
		return
	}
	id := int64(idP)
	article, err := a.ArticleUcase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleHandler) store(c *gin.Context) {
	var article domain.Article
	err := c.Bind(&article)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = a.ArticleUcase.Store(c, &article)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": article,
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case utils.ErrInternalServerError:
		return http.StatusInternalServerError
	case utils.ErrNotFound:
		return http.StatusNotFound
	case utils.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
