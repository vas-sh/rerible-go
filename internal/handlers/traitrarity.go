package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/rerible-go/internal/models"
)

func (h *handler) TraitRarity(c *gin.Context) {
	var req models.TraitRarityRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.client.TraitRarities(c.Request.Context(), req)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, res)
}
