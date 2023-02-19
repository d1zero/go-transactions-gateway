package repository

import (
	"context"
	"go-transactions-gateway/internal/domain/entity"
)

type UserRepository interface {
	Create(context.Context, entity.User) (entity.User, error)
	CreateBalance(context.Context, entity.Balance) (entity.Balance, error)
	GetUsersBalances(context.Context) (result []entity.UserBalance, err error)
}
