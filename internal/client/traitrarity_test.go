package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vas-sh/rerible-go/internal/models"
)

func TestTraitraritySuccess(t *testing.T) {
	data := models.TraitRarityRequest{}
	mockResponse := `{"traits":[{"key":"1", "rarity":"3", "value":"2"}]}`
	want := &models.TraitRarityResponse{
		Traits: []models.TraitRarity{
			{
				Key:    "1",
				Value:  "2",
				Rarity: "3",
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockResponse)
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	c := New("", mockServer.Client(), mockServer.URL)
	got, err := c.TraitRarities(context.Background(), data)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, got, want)
}

func TestTraitrarityError(t *testing.T) {
	data := models.TraitRarityRequest{}
	mockError := errors.New("not found")

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, mockError.Error(), http.StatusNotFound)
	}))
	defer mockServer.Close()

	c := New("", mockServer.Client(), mockServer.URL)
	got, err := c.TraitRarities(context.Background(), data)
	assert.Equal(t, errors.New("code: 404 Not Found, response: not found\n"), err)
	assert.Nil(t, got)
}
