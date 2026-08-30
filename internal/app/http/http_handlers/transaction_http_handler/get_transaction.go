package transaction_http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetTransaction(rw http.ResponseWriter, r *http.Request) {
	transactionID := r.PathValue("id")

	transaction, err := h.transactionUseCase.GetTransaction(r.Context(), transactionID)
	if err != nil {
		rw.Write([]byte(fmt.Errorf("transactionUseCase.GetTransaction: %w", err).Error()))

		return
	}

	data, _ := json.MarshalIndent(mapDomainTransactionToResponse(transaction), "", "  ")
	rw.Write(data)
}
