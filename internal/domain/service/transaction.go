package service

import (
	"context"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/repository"
)

type Transaction struct {
	repos repository.Registry
}

func (s *Transaction) GetTransactions(ctx context.Context, p dto.PaginationRequest) (result dto.GetTransactionsResponse, err error) {
	result.Data, err = s.repos.Transaction().Get(ctx, p)
	if err != nil {
		return dto.GetTransactionsResponse{}, err
	}

	result.Count, err = s.repos.Transaction().CountTransactions(ctx)
	if err != nil {
		return dto.GetTransactionsResponse{}, err
	}
	return result, nil
}

var _ domain.TransactionService = &Transaction{}

func NewTransactionService(repos repository.Registry) *Transaction {
	return &Transaction{repos}
}
