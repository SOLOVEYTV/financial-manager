package transaction_usecase

import (
	"context"
	"fmt"

	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
)

func (u *UseCase) DeleteTransaction(ctx context.Context, transactionID string) (*model.Transaction, error) {
	transaction, err := u.transactionRepo.DeleteTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, fmt.Errorf("transactionRepo.DeleteTransactionByID: %w", err)
	}

	return transaction, nil
}
