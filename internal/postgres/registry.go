package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"go-transactions-gateway/internal/domain/entity"
	"go-transactions-gateway/internal/domain/repository"
)

type (
	PGRegistry struct {
		db *sqlx.DB
		m  *pgManager
	}
)

func (r *PGRegistry) User() repository.UserRepository {
	return r.m.User()
}

func (r *PGRegistry) Transaction() repository.TransactionsRepository {
	return r.m.Transaction()
}

func (r *PGRegistry) WithTx(ctx context.Context, f func(repository.EntityManager) error) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = f(&pgManager{
		user:        &User{tx},
		transaction: &Transaction{tx},
	})
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.SerializationFailure {
			return entity.ErrConcurrentTx
		}
		return err
	}
	return nil
}

var _ repository.Registry = &PGRegistry{}

func NewPGRegistry(db *sqlx.DB) *PGRegistry {
	return &PGRegistry{
		db: db,
		m: &pgManager{
			user:        &User{db},
			transaction: &Transaction{db},
		},
	}
}
