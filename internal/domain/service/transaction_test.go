package service

import (
	"go-transactions-gateway/internal/domain/entity"
	"testing"
)

func TestTransaction_GetTransactions(t *testing.T) {
	tests := []struct {
		name   string
		userID int
		errMsg error
	}{
		{
			"user with transactions",
			1,
			nil,
		},
		{
			"user without transactions",
			2,
			entity.ErrNoUserTransactions,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			transactionService := NewTransactionService()
			_, err := transactionService.GetTransactions(tc.userID)

			if err != tc.errMsg {
				t.Error("error message: expected", tc.errMsg, "received", err)
			}
		})
	}
}
