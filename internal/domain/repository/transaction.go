package repository

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
)

type TransactionsRepository interface {
	FindByUserID(context.Context, dto.GetTransactionsRequest) ([]entity.Transaction, error)
}
