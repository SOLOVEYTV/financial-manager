package transaction_http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SOLOVEYTV/financial-manager/internal/service/usecases/transacation_usecase"
)

type createTransactionRequest struct {
	Type        string     `json:"type"`
	Amount      int64      `json:"amount"` // копейки
	Description *string    `json:"description"`
	Category    string     `json:"category"`
	Date        *time.Time `json:"date"`
}

func (h *Handler) CreateTransaction(rw http.ResponseWriter, r *http.Request) {
	// объявление переменной, куда будет декодирован запрос из json
	var req createTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.Write([]byte(fmt.Errorf("could not decode request body").Error()))

		return
	}

	transaction, err := h.transactionUseCase.CreateTransaction(r.Context(), transacation_usecase.CreateTransactionReq{
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Description,
		Category:    req.Category,
		Date:        req.Date,
	})
	if err != nil {
		rw.Write([]byte(fmt.Errorf("transactionUseCase.CreateTransaction").Error()))

		return
	}

	data, _ := json.MarshalIndent(mapDomainTransactionToResponse(transaction), "", "  ")
	rw.Write(data)
}
