package main

import (
	"fmt"
	"net/http"

	"github.com/SOLOVEYTV/financial-manager/cmd/service_provider"
)

func main() {
	fmt.Println("старт приложения financial-manager")

	sp := service_provider.NewServiceProvider()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /transactions", sp.GetTransactionHTTPHandler().CreateTransaction)
	mux.HandleFunc("GET /transactions/{id}", sp.GetTransactionHTTPHandler().GetTransaction)
	mux.HandleFunc("DELETE /transactions/{id}", sp.GetTransactionHTTPHandler().DeleteTransaction)

	fmt.Println("http сервер успешно запущен и готов принимать входящие запросы")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
