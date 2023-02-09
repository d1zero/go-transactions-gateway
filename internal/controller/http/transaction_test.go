package http

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/service/mocks"
	"go-transactions-gateway/pkg/govalidator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionController_GetTransactions(t *testing.T) {
	tests := []struct {
		name         string
		route        string
		data         dto.GetTransactionsRequest
		result       []entity.Transaction
		responseCode int
		errMsg       error
	}{
		{
			"valid data",
			"/api/transaction",
			dto.GetTransactionsRequest{UserID: 1},
			[]entity.Transaction{{ID: 1}},
			200,
			nil,
		},
		{
			"user without transactions data",
			"/api/transaction",
			dto.GetTransactionsRequest{UserID: 2},
			[]entity.Transaction{},
			204,
			nil,
		},
		{
			"non valid data",
			"/api/transaction",
			dto.GetTransactionsRequest{UserID: 0},
			[]entity.Transaction{},
			400,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := govalidator.New()
			app := fiber.New()
			apiGroup := app.Group("api")
			transactionGroup := apiGroup.Group("transaction")

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			body, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatal("invalid body")
			}

			req := httptest.NewRequest(http.MethodGet, tt.route, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			transactionService := mocks.NewMockTransactionService(ctrl)
			transactionController := NewTransactionService(transactionService, *v)
			transactionService.EXPECT().GetTransactions(req.Context(), tt.data).Return(tt.result, tt.errMsg).MaxTimes(1)

			transactionController.RegisterTransactionRoutes(transactionGroup)

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatal("unable to make request")
			}
			assert.Equal(t, tt.responseCode, resp.StatusCode)
		})
	}
}
