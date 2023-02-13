package service

import (
	"context"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type Transaction struct {
	repos repository.Registry
}

func (s *Transaction) GetTransactions(ctx context.Context, p dto.GetTransactionsRequest) ([]entity.Transaction, error) {
	return s.repos.Transaction().FindByUserID(ctx, p)
}

var _ domain.TransactionService = &Transaction{}

func NewTransactionService(repos repository.Registry) *Transaction {
	return &Transaction{repos}
}
