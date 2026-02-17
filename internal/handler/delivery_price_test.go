package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"golang-api-practice/internal/handler"
	"golang-api-practice/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockService struct {
	resp *model.Response
	err  error
}

func (m *mockService) CalculateTotalFee(_ context.Context, _ *model.Request) (*model.Response, error) {
	return m.resp, m.err
}

func TestDeliveryPriceHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		mockResp       *model.Response
		mockErr        error
		wantStatusCode int
		wantBody       *model.Response
	}{
		{
			name:           "success",
			query:          "?venue_slug=test-venue&cart_value=1000&user_lat=60.17&user_lon=24.93",
			mockResp:       &model.Response{TotalPrice: 1200},
			mockErr:        nil,
			wantStatusCode: http.StatusOK,
			wantBody:       &model.Response{TotalPrice: 1200},
		},
		{
			name:           "missing parameter",
			query:          "?cart_value=1000&user_lat=60.17&user_lon=24.93",
			mockResp:       nil,
			mockErr:        nil,
			wantStatusCode: http.StatusBadRequest,
			wantBody:       nil,
		},
		{
			name:           "service error",
			query:          "?venue_slug=test-venue&cart_value=1000&user_lat=60.17&user_lon=24.93",
			mockResp:       nil,
			mockErr:        errors.New("service failed"),
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup handler with mock service
			h := handler.NewHandler(&mockService{resp: tt.mockResp, err: tt.mockErr})

			req := httptest.NewRequest(http.MethodGet, "/api/v1/delivery-order-price"+tt.query, nil)
			w := httptest.NewRecorder()

			h.DeliveryPrice(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			require.Equal(t, tt.wantStatusCode, resp.StatusCode)

			if tt.wantBody != nil {
				var got model.Response
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))
				require.Equal(t, tt.wantBody.TotalPrice, got.TotalPrice)
			}
		})
	}
}
