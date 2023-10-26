package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/service"
	mock_service "github.com/KrizzMU/delivery-service/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestHandler_GetById(t *testing.T) {
	type mockBehavior func(r *mock_service.MockOrder, id string)

	tests := []struct {
		name                 string
		param                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "OK",
			param: "asd12as",
			mockBehavior: func(r *mock_service.MockOrder, id string) {
				r.EXPECT().Get(id).Return(core.Order{
					OrderUID:    "asd12as",
					TrackNumber: "testtrack",
					Locale:      "EN",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"order_uid":"asd12as","track_number":"testtrack","entry":"","Delivery":{"Name":"","Phone":"","Zip":"","City":"","Address":"","Region":"","Email":""},"Payment":{"Transaction":"","request_id":"","Currency":"","Provider":"","Amount":0,"payment_dt":0,"Bank":"","delivery_cost":0,"goods_total":0,"custom_fee":0},"Items":null,"locale":"EN","internal_signature":"","customer_id":"","delivery_service":"","shardkey":"","sm_id":0,"date_created":"","oof_shard":""}`,
		},
		{
			name:  "Error",
			param: "asd12as",
			mockBehavior: func(r *mock_service.MockOrder, id string) {
				r.EXPECT().Get(id).Return(core.Order{}, errors.New("something wrong, bro"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something wrong, bro"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockOrder(c)
			test.mockBehavior(repo, test.param)

			services := &service.Service{Order: repo}
			handler := Handler{services: services}

			r := gin.New()
			r.Handle("GET", "/order/:id", handler.GetById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order/"+test.param, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
