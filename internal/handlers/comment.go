package handlers

import (
	"net/http"
	"strconv"

	"github.com/MerBerd/blog-app/internal/models"
	"github.com/gin-gonic/gin"
)

type getAllCommentsResponse struct {
	Data []models.Comment `json:"data"`
}

func (h *Handler) getAllComment(c *gin.Context) {

	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comments, err := h.services.Comment.GetAll(articleId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})

}

func (h *Handler) createComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Comment

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Comment.Create(userId, articleId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getCommentById(c *gin.Context) {

}

func (h *Handler) changeComment(c *gin.Context) {

}

func (h *Handler) deleteComment(c *gin.Context) {

}
