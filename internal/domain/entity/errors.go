package entity

import (
	"errors"
	"fmt"
)

var (
	ErrNoUserTransactions = errors.New("user has no transactions")
	ErrConcurrentTx       = fmt.Errorf("concurrent tx")
)
