package controllers

import (
	"net/http"
	"strconv"

	"final-project-golang-individu/models"
	"final-project-golang-individu/services"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	newsService services.NewsService
}

func NewNewsController(newsService services.NewsService) *NewsController {
	return &NewsController{newsService: newsService}
}

// CreateNews godoc
// @Summary Create a new news
// @Description Create a new news item
// @Tags news
// @Accept  json
// @Produce  json
// @Param news body models.News true "News"
// @Success 201 {object} models.News
// @Router /news [post]
func (ctrl *NewsController) CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	news.UserID = userID.(uint)

	if err := ctrl.newsService.CreateNews(&news); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

// UpdateNews godoc
// @Summary Update an existing news
// @Description Update an existing news item
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path int true "News ID"
// @Param news body models.News true "News"
// @Success 200 {object} models.News
// @Router /news/{id} [put]
func (ctrl *NewsController) UpdateNews(c *gin.Context) {
	var news models.News
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	existingNews, err := ctrl.newsService.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	news.ID = existingNews.ID

	if err := ctrl.newsService.UpdateNews(&news); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

// DeleteNews godoc
// @Summary Delete a news item
// @Description Delete a news item
// @Tags news
// @Param id path int true "News ID"
// @Success 204
// @Router /news/{id} [delete]
func (ctrl *NewsController) DeleteNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	if err := ctrl.newsService.DeleteNews(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetNews godoc
// @Summary Get a single news item by ID
// @Description Get a single news item by ID
// @Tags news
// @Param id path int true "News ID"
// @Success 200 {object} models.News
// @Router /news/{id} [get]
func (ctrl *NewsController) GetNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	news, err := ctrl.newsService.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// GetAllNews godoc
// @Summary Get all news items
// @Description Get all news items
// @Tags news
// @Success 200 {array} models.News
// @Router /news [get]
func (ctrl *NewsController) GetAllNews(c *gin.Context) {
	news, err := ctrl.newsService.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}
