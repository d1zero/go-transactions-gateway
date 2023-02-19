package repository

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
)

type TransactionsRepository interface {
	FindByUserID(context.Context, int) ([]entity.Transaction, error)
	Get(context.Context, dto.PaginationRequest) ([]entity.Transaction, error)
	CountTransactions(context.Context) (int, error)
	CreateLedgerEntry(context.Context, entity.Ledger) (entity.Ledger, error)
}
