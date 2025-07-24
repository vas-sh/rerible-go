package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Ownership(c *gin.Context) {
	id := c.Param("id")
	ownership, err := h.client.Ownership(c.Request.Context(), id)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, ownership)
}
