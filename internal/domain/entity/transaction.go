package entity

import "github.com/shopspring/decimal"

type (
	Transaction struct {
		ID                string          `db:"id" json:"id"`
		ClientID          int             `db:"client_id" json:"clientID"`
		Type              string          `db:"type" json:"type"`
		Amount            decimal.Decimal `db:"amount" json:"amount"`
		CommissionFix     decimal.Decimal `db:"commission_fix" json:"commissionFix"`
		CommissionPercent decimal.Decimal `db:"commission_percent" json:"commissionPercent"`
	}
)
