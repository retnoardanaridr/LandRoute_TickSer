package routes

import (
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func TicketRoutes(r *mux.Router) {
	TicketRepository := repositories.RepositoryTicket(mysql.DB)
	h := handlers.HandlerTicket(TicketRepository)

	r.HandleFunc("/tickets", h.FindTicket).Methods("GET")
	r.HandleFunc("/ticket/{id}", h.GetTicket).Methods("GET")
	r.HandleFunc("/ticket", middleware.Auth(h.AddTicket)).Methods("POST")

	r.HandleFunc("/filter", h.FindFilter).Methods("POST")
	r.HandleFunc("/transaction-qty/{id}", middleware.Auth(h.CreateTransactionQtyTicket)).Methods("POST")
}
