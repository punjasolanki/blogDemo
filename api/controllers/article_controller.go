package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin.com/gin/models"
)

/*
*Author: Solanki
*Date: 30-Jun-2023
*Method: GET
*Description: FindArticles function get article list on desc orders array return.
 */

func FindArticles(c *gin.Context) {
	var articles []models.Article
	models.DB.Order("id DESC").Find(&articles)
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

/*
*Author: Solanki
*Date: 30-Jun-2023
*Method: GET
*params: id (article table)
*Description: FindParticularArticles function get particular article obj return.
 */
func FindParticularArticles(c *gin.Context) {
	var article models.Article
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": article})

}
