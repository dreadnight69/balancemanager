package http

import "time"

type UserInfoResponse struct {
	UserID  int64  `json:"user_id"`
	Balance string `json:"balance"`
}

type SendFundsResponse struct {
	SenderID      int64  `json:"sender_id"`
	SenderBalance string `json:"sender_balance"`
	RecipientID   int64  `json:"recipient_id"`
	Amount        string `json:"amount"`
}

type GetTransactionsResponse struct {
	Transactions []Transaction
	NextCursor   int64 `json:"next_cursor"`
}

type Transaction struct {
	TransactionID int64     `json:"transaction_id"`
	InitiatorID   int64     `json:"initiator_id"`
	RecipientID   int64     `json:"recipient_id"`
	Amount        string    `json:"amount"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
	OperationType string    `json:"operation_type"`
}

type UpdateBalanceRequest struct {
	UserID      int64  `json:"user_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

type SendFundsRequest struct {
	SenderID    int64  `json:"sender_id"`
	RecipientID int64  `json:"recipient_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}
