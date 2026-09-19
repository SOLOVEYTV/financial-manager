package service_provider

import (
	"github.com/SOLOVEYTV/financial-manager/internal/adapter/storage/in_memory_storage/transaction_in_memory_repo"
	"github.com/SOLOVEYTV/financial-manager/internal/app/http/http_handlers/transaction_http_handler"
	"github.com/SOLOVEYTV/financial-manager/internal/service/usecases/transaction_usecase"
)

type ServiceProvider struct {
	transactionInMemoryRepo *transaction_in_memory_repo.Repo

	transactionUseCase *transaction_usecase.UseCase

	transactionHTTPHandler *transaction_http_handler.Handler
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) getTransactionInMemoryRepo() *transaction_in_memory_repo.Repo {
	if s.transactionInMemoryRepo == nil {
		s.transactionInMemoryRepo = transaction_in_memory_repo.NewRepo()
	}

	return s.transactionInMemoryRepo
}

func (s *ServiceProvider) getTransactionUseCase() *transaction_usecase.UseCase {
	if s.transactionUseCase == nil {
		s.transactionUseCase = transaction_usecase.NewUseCase(
			s.getTransactionInMemoryRepo(),
		)
	}

	return s.transactionUseCase
}

func (s *ServiceProvider) GetTransactionHTTPHandler() *transaction_http_handler.Handler {
	if s.transactionHTTPHandler == nil {
		s.transactionHTTPHandler = transaction_http_handler.NewHandler(
			s.getTransactionUseCase(),
		)
	}

	return s.transactionHTTPHandler
}
