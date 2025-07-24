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

func TestOwnershipSuccess(t *testing.T) {
	id := "123"
	mockResponse := `{"blockchain":"", "contract":"", "id":"", "owner":"0x123", "tokenId":"", "value":""}`
	want := &models.Ownership{
		Owner: "0x123",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockResponse)
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	c := New("", mockServer.Client(), mockServer.URL)
	got, err := c.Ownership(context.Background(), id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, got, want)
}

func TestOwnershipError(t *testing.T) {
	id := "not_found"
	mockError := errors.New("not found")

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, mockError.Error(), http.StatusNotFound)
	}))
	defer mockServer.Close()

	c := New("", mockServer.Client(), mockServer.URL)
	got, err := c.Ownership(context.Background(), id)
	assert.Equal(t, errors.New("code: 404 Not Found, response: not found\n"), err)
	assert.Nil(t, got)
}
