package dto

type (
	PaginationRequest struct {
		Limit  *int `query:"limit" validate:"required"`
		Offset *int `query:"offset" validate:"required"`
	}

	PaginationResponse struct {
		Count int `json:"count"`
	}
)
