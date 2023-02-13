package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type Transaction struct {
	q queryRunner
}

func (r *Transaction) FindByUserID(ctx context.Context,
	p dto.GetTransactionsRequest) (result []entity.Transaction, err error) {
	return []entity.Transaction{}, nil
}

var _ repository.TransactionsRepository = &Transaction{}

func NewTransactionRepository(db *sqlx.DB) *Transaction {
	return &Transaction{db}
}
