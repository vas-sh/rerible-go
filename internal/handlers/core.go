package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/rerible-go/internal/models"
)

//go:generate mockgen -source=core.go -destination=mocks/mocks.go -package mocks

type reribler interface {
	Ownership(ctx context.Context, id string) (*models.Ownership, error)
	TraitRarities(ctx context.Context, data models.TraitRarityRequest) (*models.TraitRarityResponse, error)
}

type handler struct {
	client reribler
}

func New(client reribler) *handler {
	return &handler{
		client: client,
	}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/ownership/:id", h.Ownership)
	r.POST("/rarity", h.TraitRarity)
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}
