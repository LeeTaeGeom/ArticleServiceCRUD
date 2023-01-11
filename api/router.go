package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/articles/:id", getArticle)
	r.GET("/api/v1/articles", getArticles)
	r.POST("/api/v1/articles", postArticle)
	r.PUT("/api/v1/articles/:id", putArticle)
	r.DELETE("/api/v1/articles/:id", deleteArticle)
	return r
}
