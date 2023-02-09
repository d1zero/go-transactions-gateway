package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/postgres/mocks"
	"testing"
)

func TestTransaction_GetTransactions(t *testing.T) {
	tests := []struct {
		name   string
		data   dto.GetTransactionsRequest
		errMsg error
	}{
		{
			"user with transactions",
			dto.GetTransactionsRequest{UserID: 1},
			nil,
		},
		{
			"user without transactions",
			dto.GetTransactionsRequest{UserID: 2},
			entity.ErrNoUserTransactions,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			transactionRepository := mocks.NewMockTransactionsRepository(ctrl)
			transactionService := NewTransactionService(transactionRepository)
			transactionRepository.EXPECT().FindByUserID(context.Background(), tc.data).MaxTimes(1)

			_, err := transactionService.GetTransactions(context.Background(), tc.data)

			if err != tc.errMsg {
				t.Error("error message: expected", tc.errMsg, "received", err)
			}
		})
	}
}
