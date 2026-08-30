package transaction_http_handler

import (
	"time"

	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
)

type transactionResponse struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	Description *string   `json:"description,omitempty"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

func mapDomainTransactionToResponse(tr *model.Transaction) transactionResponse {
	return transactionResponse{
		ID:          tr.ID,
		CreatedAt:   tr.CreatedAt,
		UpdatedAt:   tr.UpdatedAt,
		DeletedAt:   tr.DeletedAt,
		Type:        tr.Type,
		Amount:      tr.Amount,
		Description: tr.Description,
		Category:    tr.Category,
		Date:        tr.Date,
	}
}
