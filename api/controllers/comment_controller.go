package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin.com/gin/models"
)

type CreateCommentInput struct {
	ArticleId uint   `json:"article_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

/*
*Author: Solanki
*Date: 30-Jun-2023
*Method: GET
*params: article_id (article table)
*Description: FindCommentsByArticle function get particular article all comment on desc orders.
 */
func FindCommentsByArticle(c *gin.Context) {
	var comments []models.Comment
	articleId := c.Param("article_id")
	if err := models.DB.Where("article_id = ?", articleId).Where("is_deleted = ?", 0).Order("id DESC").Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comments not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

/*
*Author: Solanki
*Date: 30-Jun-2023
*Method: POST
*Body: articleId,name,content
*Description: AddComment function pass Json obj to store comments table.
 */
func AddComment(c *gin.Context) {
	// Validate input
	var input CreateCommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	commentData := models.Comment{ArticleId: input.ArticleId, Name: input.Name, Content: input.Content}
	models.DB.Create(&commentData)

	c.JSON(http.StatusCreated, gin.H{"data": commentData})

}

/*
*Author: Solanki
*Date: 30-Jun-2023
*Method: PATCH
*params: comment_id
*Description: RemoveComment function soft deleted particular comments.
 */
func RemoveComment(c *gin.Context) {
	var comment models.Comment

	commentId := c.Param("comment_id")

	if err := models.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	models.DB.Model(&comment).Where("id = ?", commentId).Update("is_deleted", true)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}
