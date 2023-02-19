package domain

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
)

type TransactionService interface {
	GetTransactions(context.Context, dto.PaginationRequest) (dto.GetTransactionsResponse, error)
}
