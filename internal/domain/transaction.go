package domain

import "go-transactions-gateway/internal/domain/entity"

type TransactionService interface {
	GetTransactions(int) ([]entity.Transaction, error)
}
