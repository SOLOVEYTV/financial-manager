package transacation_usecase

import (
	"github.com/SOLOVEYTV/financial-manager/internal/adapter/storage/in_memory_storage/transaction_in_memory_repo"
)

type UseCase struct {
	transactionRepo *transaction_in_memory_repo.Repo
}

func NewUseCase(transactionRepo *transaction_in_memory_repo.Repo) *UseCase {
	return &UseCase{
		transactionRepo: transactionRepo,
	}
}
