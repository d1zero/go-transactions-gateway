package postgres

import (
	"context"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type Transaction struct {
	q queryRunner
}

func (r *Transaction) FindByUserID(ctx context.Context,
	userID int) (result []entity.Transaction, err error) {
	q := `
		SELECT 
			id, 
			client_id, 
			type, 
			amount, 
			commission_fix, 
			commission_percent
		FROM clients.transactions 
		WHERE client_id=$1;
	`

	err = r.q.SelectContext(ctx, &result, q, userID)

	if err != nil {
		return []entity.Transaction{}, err
	}

	return result, nil
}

func (r *Transaction) CreateLedgerEntry(ctx context.Context, p entity.Ledger) (result entity.Ledger, err error) {
	q := `
		INSERT INTO clients.ledger
			(
			 account_id, 
			 base_amount, 
			 action_amount, 
			 action_type
			) 
		VALUES 
		    (
		     $1,
		     $2,
		     $3,
		     $4
			)
		RETURNING 
		    id, 
		    account_id, 
		    base_amount, 
		    action_amount, 
		    action_type
		;
	`

	err = r.q.GetContext(ctx, &result, q, p.AccountID, p.BaseAmount, p.ActionAmount, p.ActionType)
	if err != nil {
		return entity.Ledger{}, err
	}

	return result, nil
}

func (r *Transaction) Get(ctx context.Context, p dto.PaginationRequest) (result []entity.Transaction, err error) {
	q := `
		SELECT 
			id, 
			client_id, 
			type, 
			amount, 
			commission_fix, 
			commission_percent
		FROM clients.transactions 
		LIMIT $1 OFFSET $2;
	`

	err = r.q.SelectContext(ctx, &result, q, p.Limit, p.Offset)

	if err != nil {
		return []entity.Transaction{}, err
	}

	return result, nil
}

func (r *Transaction) CountTransactions(ctx context.Context) (count int, err error) {
	q := `
		SELECT COUNT(*) FROM clients.transactions 
	`
	err = r.q.GetContext(ctx, &count, q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

var _ repository.TransactionsRepository = &Transaction{}
