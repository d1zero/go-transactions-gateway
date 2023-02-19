package entity

import "github.com/shopspring/decimal"

type (
	User struct {
		ID                int64           `db:"id" json:"id"`
		FirstName         string          `db:"first_name" json:"firstName"`
		LastName          string          `db:"last_name" json:"lastName"`
		CommissionFix     decimal.Decimal `db:"commission_fix" json:"commissionFix"`
		CommissionPercent decimal.Decimal `db:"commission_percent" json:"commissionPercent"`
	}

	Balance struct {
		ID          string `db:"id" json:"id"`
		ClientID    int64  `db:"client_id" json:"clientID"`
		Description string `db:"description" json:"description"`
	}

	Ledger struct {
		ID           string          `db:"id" json:"id"`
		AccountID    string          `db:"account_id" json:"accountID"`
		BaseAmount   decimal.Decimal `db:"base_amount" json:"baseAmount"`
		ActionAmount decimal.Decimal `db:"action_amount" json:"actionAmount"`
		ActionType   string          `db:"action_type" json:"actionType"`
	}

	UserBalance struct {
		ID       string          `db:"id" json:"id"`
		ClientID int64           `db:"client_id" json:"clientID"`
		Client   string          `db:"client" json:"client"`
		Amount   decimal.Decimal `db:"balance" json:"balance"`
	}
)
