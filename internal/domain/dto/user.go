package dto

type (
	CreateUser struct {
		FirstName         string `json:"firstName" validate:"required"`
		LastName          string `json:"lastName" validate:"required"`
		CommissionFix     string `json:"commissionFix" validate:"required"`
		CommissionPercent string `json:"commissionPercent" validate:"required"`
	}

	GetUserTransactionsRequest struct {
		UserID int `params:"user_id" validate:"required"`
	}
)
