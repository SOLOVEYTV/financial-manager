package transaction_http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	"github.com/SOLOVEYTV/financial-manager/internal/app/http/http_helpers"
	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
	"github.com/SOLOVEYTV/financial-manager/internal/service/usecases/transaction_usecase"
)

type createTransactionRequest struct {
	Type        string     `json:"type"`
	Amount      int64      `json:"amount"` // копейки
	Description *string    `json:"description"`
	Category    string     `json:"category"`
	Date        *time.Time `json:"date"`
}

func (h *Handler) CreateTransaction(rw http.ResponseWriter, r *http.Request) {
	rh := http_helpers.NewResponseHandler(rw)

	// объявление переменной, куда будет декодирован запрос из json
	var req createTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rh.ErrorResponse(http.StatusBadRequest, err.Error(), http_helpers.CodeDecodeFailed)

		return
	}

	validatedUcReq, validationErr := validateCreateTransactionRequest(req)
	if validationErr != nil {
		rh.ErrorResponse(
			http.StatusBadRequest,
			"invalid request",
			validationErr.Error(),
			validationErr.Violations()...,
		)

		return
	}

	transaction, err := h.transactionUseCase.CreateTransaction(r.Context(), validatedUcReq)
	if err != nil {
		rh.ErrorResponse(
			http.StatusInternalServerError,
			fmt.Errorf("transactionUseCase.CreateTransaction: %w", err).Error(),
			http_helpers.CodeInternalError,
		)

		return
	}

	rh.SuccessResponse(http.StatusCreated, mapDomainTransactionToResponse(transaction))
}

func validateCreateTransactionRequest(req createTransactionRequest) (
	transaction_usecase.CreateTransactionReq, *http_helpers.ValidationError,
) {
	err := http_helpers.NewValidationError()

	if req.Amount <= 0 {
		err = err.Add("amount", http_helpers.ErrTxtMustBePositive)
	}
	if req.Description != nil {
		if *req.Description == "" {
			err = err.Add("description", http_helpers.ErrTxtEmpty)
		}
		const maxDescriptionLength = 256
		if utf8.RuneCountInString(*req.Description) > maxDescriptionLength {
			err = err.Add("description", http_helpers.ErrTxtTooLong)
		}
	}
	if req.Date != nil && req.Date.IsZero() {
		err = err.Add("date", http_helpers.ErrTxtZeroValue)
	}

	transactionType := parseTransactionType(req.Type)
	if transactionType == model.TransactionTypeUnspecified {
		err = err.Add("type", http_helpers.ErrTxtUnknown)
	}

	transactionCategory := parseTransactionCategory(req.Category)
	if transactionCategory == model.TransactionCategoryUnspecified {
		err = err.Add("category", http_helpers.ErrTxtUnknown)
	}

	if len(err.Violations()) > 0 {
		return transaction_usecase.CreateTransactionReq{}, err
	}

	return transaction_usecase.CreateTransactionReq{
		Type:        transactionType,
		Amount:      req.Amount,
		Description: req.Description,
		Category:    transactionCategory,
		Date:        req.Date,
	}, nil
}
