package postgres

import (
	"context"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type User struct {
	q queryRunner
}

func (r *User) Create(ctx context.Context, p entity.User) (result entity.User, err error) {
	q := `
		INSERT INTO clients.users 
		    (first_name, last_name, commission_fix, commission_percent)
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, commission_fix, commission_percent;
	`
	err = r.q.GetContext(ctx, &result, q, p.FirstName, p.LastName, p.CommissionFix, p.CommissionPercent)
	return result, err
}

func (r *User) CreateBalance(ctx context.Context, p entity.Balance) (result entity.Balance, err error) {
	q := `
		INSERT INTO clients.balance
			(client_id, description) 
		VALUES ($1, $2)
		RETURNING id, client_id, description;
	`
	err = r.q.GetContext(ctx, &result, q, p.ClientID, p.Description)
	return result, err
}

func (r *User) GetUsersBalances(ctx context.Context) (result []entity.UserBalance, err error) {
	q := `
		SELECT 
		    b.id,
		    u.id AS client_id,
		    CONCAT(u.first_name, ' ', u.last_name) AS client,
		    l.base_amount+l.action_amount AS balance 
		FROM clients.users u
		INNER JOIN clients.balance b ON b.client_id=u.id
		INNER JOIN clients.ledger l ON l.account_id=b.id
	`
	err = r.q.SelectContext(ctx, &result, q)
	return result, err
}

var _ repository.UserRepository = &User{}
