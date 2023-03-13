package transactiondto

type TransactionRequest struct {
	UserID   int    `json:"user_id"`
	TicketID int    `json:"ticket_id"`
	Total    int    `json:"total"`
	Qty      int    `json:"qty"`
	Status   string `json:"status"`
}
