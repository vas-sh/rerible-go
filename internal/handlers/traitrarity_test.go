package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vas-sh/rerible-go/internal/handlers/mocks"
	"github.com/vas-sh/rerible-go/internal/models"
	"go.uber.org/mock/gomock"
)

func TestTraitRarity(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		mockResponse   *models.TraitRarityResponse
		mockError      error
		expectedStatus int
		want           string
	}{
		{
			name: "ok",
			body: strings.NewReader(`{"collectionId": "id","properties": [{"key": "Hat","value": "Halo"}]}`),
			mockResponse: &models.TraitRarityResponse{
				Traits: []models.TraitRarity{
					{
						Key:    "key1",
						Value:  "3",
						Rarity: "5",
					},
					{
						Key:    "key2",
						Value:  "4",
						Rarity: "2",
					},
				},
			},
			expectedStatus: http.StatusOK,
			want:           `{"traits":[{"key":"key1", "rarity":"5", "value":"3"}, {"key":"key2", "rarity":"2", "value":"4"}]}`,
		},
		{
			name:           "error",
			body:           strings.NewReader(`{"collectionId": "id","properties": [{"key": "Hat","value": "Halo"}]}`),
			mockError:      errors.New("not found"),
			expectedStatus: http.StatusBadRequest,
			want:           "not found\n",
		},
		{
			name:           "no body",
			body:           http.NoBody,
			expectedStatus: http.StatusBadRequest,
			want:           "EOF\n",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			raribleMock := mocks.NewMockreribler(gomock.NewController(t))
			if ts.mockResponse != nil || ts.mockError != nil {
				raribleMock.EXPECT().TraitRarities(gomock.Any(), gomock.Any()).Return(ts.mockResponse, ts.mockError)
			}
			h := New(raribleMock)
			router := gin.Default()
			h.Register(router)

			req, err := http.NewRequest(http.MethodPost, "/rarity", ts.body)
			if err != nil {
				t.Error(err)
				return
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, ts.expectedStatus, w.Code)
			if w.Code == http.StatusOK {
				assert.JSONEq(t, ts.want, w.Body.String())
			} else {
				assert.Equal(t, ts.want, w.Body.String())
			}
		})
	}
}
