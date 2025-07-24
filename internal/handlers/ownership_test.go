package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vas-sh/rerible-go/internal/handlers/mocks"
	"github.com/vas-sh/rerible-go/internal/models"
	"go.uber.org/mock/gomock"
)

func TestOwnership(t *testing.T) {
	testCases := []struct {
		name           string
		id             string
		mockResponse   *models.Ownership
		mockError      error
		expectedStatus int
		want           string
	}{
		{
			name: "ok",
			id:   "123",
			mockResponse: &models.Ownership{
				Owner: "0x123",
			},
			expectedStatus: http.StatusOK,
			want:           `{"blockchain":"", "contract":"", "id":"", "owner":"0x123", "tokenId":"", "value":""}`,
		},
		{
			name:           "error",
			id:             "not_found",
			mockError:      errors.New("not found"),
			expectedStatus: http.StatusBadRequest,
			want:           "not found\n",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			raribleMock := mocks.NewMockreribler(gomock.NewController(t))
			raribleMock.EXPECT().Ownership(gomock.Any(), ts.id).Return(ts.mockResponse, ts.mockError)
			h := New(raribleMock)
			router := gin.Default()
			h.Register(router)

			req, err := http.NewRequest(http.MethodGet, "/ownership/"+ts.id, nil)
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
