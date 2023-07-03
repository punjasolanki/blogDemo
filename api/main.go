package main

import (
	"github.com/gin-gonic/gin"

	"gin.com/gin/controllers"
	"gin.com/gin/models"
)

func main() {
	route := gin.Default()
	models.ConnectDatabase()
	route.Use(CORSMiddleware())

	route.GET("/api/articles", controllers.FindArticles)
	route.GET("/api/article/:id", controllers.FindParticularArticles)
	route.GET("/api/article/comments/:article_id", controllers.FindCommentsByArticle)
	route.POST("/api/article/comment", controllers.AddComment)
	route.PATCH("/api/article/comment/:comment_id", controllers.RemoveComment)
	route.Run()
}

//Access control Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
