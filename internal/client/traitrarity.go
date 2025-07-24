package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/vas-sh/rerible-go/internal/models"
)

func (c *client) TraitRarities(ctx context.Context, data models.TraitRarityRequest) (*models.TraitRarityResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var res models.TraitRarityResponse
	err = c.send(ctx, "/items/traits/rarity", http.MethodPost, bytes.NewBuffer(body), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
