package transaction_http_handler

import "github.com/SOLOVEYTV/financial-manager/internal/service/usecases/transacation_usecase"

type Handler struct {
	transactionUseCase *transacation_usecase.UseCase
}

func NewHandler(transactionUseCase *transacation_usecase.UseCase) *Handler {
	return &Handler{
		transactionUseCase: transactionUseCase,
	}
}
