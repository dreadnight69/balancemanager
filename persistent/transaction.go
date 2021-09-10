package persistent

import (
	"balanceManager/models"
	"context"
	"time"
)

type TransactionsStorage interface {
	CreateTransaction(ctx context.Context, initiatorID int64, recipientID int64, amount int64, description string, date time.Time, operationTypeID int) (transactionID int64, err error)
	GetTransactionsByUserID(ctx context.Context, userID int64, limit uint64, timeCursor time.Time) (*[]models.Transaction, error)
	CountTransactionsLeft(ctx context.Context, userID int64, timeCursor time.Time) (count int, err error)
}
