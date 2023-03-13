package routes

import (
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	r.HandleFunc("/transactions", h.FindTransaction).Methods("GET")
	r.HandleFunc("/create-trans/{id}", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/order-user", middleware.Auth(h.GetTransByUser)).Methods("GET")
	r.HandleFunc("/get-idpayment/{id}", middleware.Auth(h.GetIdPayment)).Methods("GET")
	r.HandleFunc("/payment/{id}", middleware.Auth(h.PaymentTrans)).Methods("GET")

	// r.HandleFunc("/transaction-user", middleware.Auth(h.GetTransactionUser)).Methods("GET")
	// r.HandleFunc("/payment", middleware.Auth(h.CreatePayment)).Methods("POST")
}
