package postgres

import "go-transactions-gateway/internal/domain/repository"

type (
	pgManager struct {
		user        *User
		transaction *Transaction
	}
)

func (m *pgManager) User() repository.UserRepository {
	return m.user
}

func (m *pgManager) Transaction() repository.TransactionsRepository {
	return m.transaction
}

var _ repository.EntityManager = &pgManager{}
