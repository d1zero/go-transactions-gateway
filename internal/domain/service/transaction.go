package service

import (
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/entity"
)

type Transaction struct {
}

func (s *Transaction) GetTransactions(userID int) ([]entity.Transaction, error) {
	if userID == 2 {
		return []entity.Transaction{}, entity.ErrNoUserTransactions
	}
	return []entity.Transaction{
		{ID: 1},
	}, nil
}

var _ domain.TransactionService = &Transaction{}

func NewTransactionService() *Transaction {
	return &Transaction{}
}
