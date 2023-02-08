package dto

type (
	GetTransactionsRequest struct {
		UserID int `json:"userID" validate:"required"`
	}
)
