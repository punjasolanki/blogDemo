package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gin.com/gin/controllers"
	"gin.com/gin/models"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	models.ConnectDatabase()
	return router
}

func TestFindArticles(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/articles", controllers.FindArticles)
	req, _ := http.NewRequest("GET", "/api/articles", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var articles []models.Article
	models.DB.Find(&articles)

	json.Unmarshal(w.Body.Bytes(), &articles)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, articles)
}

func TestFindParticularArticles(t *testing.T) {
	id := `2`
	r := SetUpRouter()
	r.GET("/api/article/:id", controllers.FindParticularArticles)
	req, _ := http.NewRequest("GET", "/api/article/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var article models.Article
	models.DB.Where("id = ?", id).First(&article)

	json.Unmarshal(w.Body.Bytes(), &article)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, article)
}

func TestFindCommentsByArticle(t *testing.T) {
	articleId := `1`
	r := SetUpRouter()
	r.GET("/api/article/comments/:id", controllers.FindCommentsByArticle)
	req, _ := http.NewRequest("GET", "/api/article/comments/"+articleId, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var comments []models.Comment
	models.DB.Where("article_id = ?", articleId).Where("is_deleted = ?", 0).Find(&comments)

	json.Unmarshal(w.Body.Bytes(), &comments)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, comments)
}

func TestAddComment(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/article/comment", controllers.AddComment)
	commentData := models.Comment{
		ArticleId: 1,
		Name:      "john",
		Content:   "unit test comment",
	}
	jsonValue, _ := json.Marshal(commentData)
	req, _ := http.NewRequest("POST", "/api/article/comment", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestRemoveComment(t *testing.T) {
	commentId := `2`
	r := SetUpRouter()
	
	r.PUT("/api/article/comment/:comment_id", controllers.RemoveComment)
	comment := models.Comment{
		IsDeleted: true,
	}

	jsonValue, _ := json.Marshal(comment)
	reqFound, _ := http.NewRequest("PUT", "/api/article/comment/"+commentId, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
}
