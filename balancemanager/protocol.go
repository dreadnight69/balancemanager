package balancemanager

import (
	"time"
)

type UserInfo struct {
	UserID  int64
	Balance string
}

type SendFundsResponse struct {
	SenderID      int64
	SenderBalance string
	RecipientID   int64
	Amount        string
}

type TransactionResponse struct {
	TransactionID int64
	InitiatorID   int64
	RecipientID   int64
	Amount        string
	OperationType string
	Description   string
	Date          time.Time
}
