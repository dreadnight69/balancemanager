package balancemanager

import (
	"context"
	"errors"
)

type Service interface {
	GetUserInfo(ctx context.Context, userID int64) (*UserInfo, error)
	MakeDeposit(ctx context.Context, userID int64, amount string, description string) (*UserInfo, error)
	WithdrawFunds(ctx context.Context, userID int64, amount string, description string) (*UserInfo, error)
	SendFunds(ctx context.Context, senderID int64, recipientID int64, amount string, description string) (*SendFundsResponse, error)
	GetTransactions(ctx context.Context, userID string, limit string, cursor string) (response *[]TransactionResponse, nextCursorUnix int64, err error)

	CreateUser(ctx context.Context) (*UserInfo, error)
}

var ErrPgNoRows = errors.New("pg: no rows in result set")
