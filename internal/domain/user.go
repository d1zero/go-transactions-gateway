package domain

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
)

type UserService interface {
	CreateUser(context.Context, dto.CreateUser) (entity.User, error)
}
