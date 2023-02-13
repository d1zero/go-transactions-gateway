package service

import (
	"context"
	"github.com/shopspring/decimal"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type User struct {
	repos repository.Registry
}

func (u *User) CreateUser(ctx context.Context, p dto.CreateUser) (user entity.User, err error) {
	commissionFix, err := decimal.NewFromString(p.CommissionFix)
	if err != nil {
		return entity.User{}, err
	}
	commissionPercent, err := decimal.NewFromString(p.CommissionPercent)
	if err != nil {
		return entity.User{}, err
	}

	err = u.repos.WithTx(ctx, func(m repository.EntityManager) (err error) {
		user, err := m.User().Create(ctx, entity.User{
			FirstName:         p.FirstName,
			LastName:          p.LastName,
			CommissionFix:     commissionFix,
			CommissionPercent: commissionPercent,
		})

		if err != nil {
			return err
		}

		_, err = m.User().CreateBalance(ctx, entity.Balance{
			ClientID:    user.ID,
			Description: "active",
		})

		if err != nil {
			return err
		}

		_, err = m.User().CreateBalance(ctx, entity.Balance{
			ClientID:    user.ID,
			Description: "frozen",
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

var _ domain.UserService = &User{}

func NewUserService(repos repository.Registry) *User {
	return &User{repos}
}
