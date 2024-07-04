package controllers

import (
	"net/http"
	"strconv"

	"final-project-golang-individu/models"
	"final-project-golang-individu/services"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment for a news item
// @Tags comment
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Comment"
// @Success 201 {object} models.Comment
// @Router /comments [post]
func (ctrl *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	comment.UserID = userID.(uint)

	if err := ctrl.commentService.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetComment godoc
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comment
// @Param id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func (ctrl *CommentController) GetComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	comment, err := ctrl.commentService.GetCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetCommentsByNews godoc
// @Summary Get comments by news ID
// @Description Get comments for a specific news item
// @Tags comment
// @Param news_id path int true "News ID"
// @Success 200 {array} models.Comment
// @Router /news/{news_id}/comments [get]
func (ctrl *CommentController) GetCommentsByNews(c *gin.Context) {
	newsID, err := strconv.Atoi(c.Param("news_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	comments, err := ctrl.commentService.GetCommentsByNewsID(uint(newsID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment by ID
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [put]
func (ctrl *CommentController) UpdateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ID = uint(id)

	if err := ctrl.commentService.UpdateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment by ID
// @Tags comment
// @Param id path int true "Comment ID"
// @Success 204
// @Router /comments/{id} [delete]
func (ctrl *CommentController) DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := ctrl.commentService.DeleteComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
