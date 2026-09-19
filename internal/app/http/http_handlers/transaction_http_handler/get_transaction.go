package transaction_http_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/SOLOVEYTV/financial-manager/internal/app/http/http_helpers"
	"github.com/SOLOVEYTV/financial-manager/internal/domain"
	"github.com/google/uuid"
)

func (h *Handler) GetTransaction(rw http.ResponseWriter, r *http.Request) {
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

	transaction, err := h.transactionUseCase.GetTransaction(r.Context(), transactionID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTransactionNotFound):
			rh.ErrorResponse(http.StatusNotFound, fmt.Errorf("transactionUseCase.GetTransaction: %w", err).Error(), http_helpers.CodeNotFound)
		default:
			rh.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("transactionUseCase.GetTransaction: %w", err).Error(), http_helpers.CodeInternalError)
		}

		return
	}

	rh.SuccessResponse(http.StatusOK, mapDomainTransactionToResponse(transaction))
}

func validateTransactionID(transactionID string) *http_helpers.ValidationError {
	err := http_helpers.NewValidationError()

	if uuid.Validate(transactionID) != nil {
		err = err.Add("id", http_helpers.ErrTxtInvalidFormat)
	}

	if len(err.Violations()) > 0 {
		return err
	}

	return nil
}
