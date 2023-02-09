package domain

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
)

type TransactionService interface {
	GetTransactions(context.Context, dto.GetTransactionsRequest) ([]entity.Transaction, error)
}
