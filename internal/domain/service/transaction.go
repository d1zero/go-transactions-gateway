package service

import (
	"context"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type Transaction struct {
	transactionRepo repository.TransactionsRepository
}

func (s *Transaction) GetTransactions(ctx context.Context, p dto.GetTransactionsRequest) ([]entity.Transaction, error) {
	if p.UserID == 2 {
		return []entity.Transaction{}, entity.ErrNoUserTransactions
	}
	return []entity.Transaction{
		{ID: 1},
	}, nil
}

var _ domain.TransactionService = &Transaction{}

func NewTransactionService(transactionRepo repository.TransactionsRepository) *Transaction {
	return &Transaction{transactionRepo}
}
