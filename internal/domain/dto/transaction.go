package dto

import "go-transactions-gateway/internal/domain/entity"

type (
	GetTransactionsRequest struct {
		UserID int `query:"user_id" validate:"required"`
	}

	GetTransactionsResponse struct {
		Data []entity.Transaction `json:"transactions"`
		PaginationResponse
	}
)
