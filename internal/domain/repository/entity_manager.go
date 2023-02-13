package repository

type (
	EntityManager interface {
		User() UserRepository
		Transaction() TransactionsRepository
	}
)
