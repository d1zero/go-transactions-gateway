package entity

import "github.com/shopspring/decimal"

type (
	User struct {
		ID                int64           `db:"id"`
		FirstName         string          `db:"first_name"`
		LastName          string          `db:"last_name"`
		CommissionFix     decimal.Decimal `db:"commission_fix"`
		CommissionPercent decimal.Decimal `db:"commission_percent"`
	}

	Balance struct {
		ID          string `db:"id"`
		ClientID    int64  `db:"client_id"`
		Description string `db:"description"`
	}
)
