package transaction_http_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/SOLOVEYTV/financial-manager/internal/app/http/http_helpers"
	"github.com/SOLOVEYTV/financial-manager/internal/domain"
)

func (h *Handler) DeleteTransaction(rw http.ResponseWriter, r *http.Request) {
	rh := http_helpers.NewResponseHandler(rw)

	transactionID := r.PathValue("id")

	if validationErr := validateTransactionID(transactionID); validationErr != nil {
		rh.ErrorResponse(
			http.StatusBadRequest,
			"invalid request",
			validationErr.Error(),
			validationErr.Violations()...,
		)

		return
	}

	_, err := h.transactionUseCase.DeleteTransaction(r.Context(), transactionID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTransactionNotFound):
			rh.ErrorResponse(http.StatusNotFound, fmt.Errorf("transactionUseCase.DeleteTransaction: %w", err).Error(), http_helpers.CodeNotFound)
		default:
			rh.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("transactionUseCase.DeleteTransaction: %w", err).Error(), http_helpers.CodeInternalError)
		}

		return
	}

	rh.SuccessResponse(http.StatusNoContent)
}
