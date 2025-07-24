package client

import (
	"context"
	"net/http"

	"github.com/vas-sh/rerible-go/internal/models"
)

func (c *client) Ownership(ctx context.Context, id string) (*models.Ownership, error) {
	var res models.Ownership
	err := c.send(ctx, "/ownerships/"+id, http.MethodGet, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
