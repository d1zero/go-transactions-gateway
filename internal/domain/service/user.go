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

func (s *User) CreateUser(ctx context.Context, p dto.CreateUser) (user entity.User, err error) {
	commissionFix, err := decimal.NewFromString(p.CommissionFix)
	if err != nil {
		return entity.User{}, err
	}
	commissionPercent, err := decimal.NewFromString(p.CommissionPercent)
	if err != nil {
		return entity.User{}, err
	}

	err = s.repos.WithTx(ctx, func(m repository.EntityManager) (err error) {
		user, err = m.User().Create(ctx, entity.User{
			FirstName:         p.FirstName,
			LastName:          p.LastName,
			CommissionFix:     commissionFix,
			CommissionPercent: commissionPercent,
		})

		if err != nil {
			return err
		}

		activeBalance, err := m.User().CreateBalance(ctx, entity.Balance{
			ClientID:    user.ID,
			Description: "active",
		})

		if err != nil {
			return err
		}

		_, err = m.Transaction().CreateLedgerEntry(ctx, entity.Ledger{
			AccountID:    activeBalance.ID,
			BaseAmount:   decimal.NewFromInt(0),
			ActionAmount: decimal.NewFromInt(0),
			ActionType:   "init",
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

func (s *User) GetUserTransactions(ctx context.Context, p dto.GetUserTransactionsRequest) ([]entity.Transaction, error) {
	return s.repos.Transaction().FindByUserID(ctx, p.UserID)
}

func (s *User) GetUsersBalances(ctx context.Context) ([]entity.UserBalance, error) {
	return s.repos.User().GetUsersBalances(ctx)
}

var _ domain.UserService = &User{}

func NewUserService(repos repository.Registry) *User {
	return &User{repos}
}
