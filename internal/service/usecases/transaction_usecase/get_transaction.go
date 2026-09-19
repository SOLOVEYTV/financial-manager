package transaction_usecase

import (
	"context"
	"fmt"

	"github.com/SOLOVEYTV/financial-manager/internal/domain"
	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
)

func (u *UseCase) GetTransaction(ctx context.Context, transactionID string) (*model.Transaction, error) {
	transaction, err := u.transactionRepo.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, fmt.Errorf("transactionRepo.GetTransactionByID: %w", err)
	}

	if transaction.DeletedAt != nil {
		return nil, fmt.Errorf("transactionRepo.GetTransactionByID: %w", domain.ErrTransactionNotFound)
	}

	return transaction, nil
}
