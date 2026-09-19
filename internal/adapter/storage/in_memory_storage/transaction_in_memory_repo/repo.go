package transaction_in_memory_repo

import (
	"context"
	"time"

	"github.com/SOLOVEYTV/financial-manager/internal/domain"
	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
)

type Repo struct {
	storage map[string]*model.Transaction
}

func NewRepo() *Repo {
	return &Repo{
		storage: make(map[string]*model.Transaction),
	}
}

func (r *Repo) CreateTransaction(_ context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	r.storage[transaction.ID] = transaction

	return transaction, nil
}

func (r *Repo) GetTransactionByID(_ context.Context, transactionID string) (*model.Transaction, error) {
	transaction, ok := r.storage[transactionID]
	if !ok {
		return nil, domain.ErrTransactionNotFound
	}

	return transaction, nil
}

func (r *Repo) DeleteTransactionByID(_ context.Context, transactionID string) (*model.Transaction, error) {
	transaction, ok := r.storage[transactionID]
	if !ok {
		return nil, domain.ErrTransactionNotFound
	}
	now := time.Now().UTC()
	transaction.DeletedAt = &now
	transaction.UpdatedAt = now

	return transaction, nil
}
