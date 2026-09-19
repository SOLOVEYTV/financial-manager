package transaction_usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
	"github.com/google/uuid"
)

type CreateTransactionReq struct {
	Type        model.TransactionType
	Amount      int64
	Description *string
	Category    model.TransactionCategory
	Date        *time.Time
}

func (u *UseCase) CreateTransaction(ctx context.Context, req CreateTransactionReq) (*model.Transaction, error) {
	id, _ := uuid.NewV7()

	now := time.Now().UTC()

	dateTr := time.Now().UTC()
	if req.Date != nil {
		dateTr = *req.Date
	}

	transaction := model.NewTransaction(
		id.String(),
		now,
		req.Type,
		req.Amount,
		req.Description,
		req.Category,
		dateTr,
	)

	createdTransaction, err := u.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("transactionRepo.CreateTransaction: %w", err)
	}

	return createdTransaction, nil
}
