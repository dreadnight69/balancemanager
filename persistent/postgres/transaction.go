package postgres

import (
	"balanceManager/models"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"time"
)

type transactionsPostgres struct {
	baseModelPostgres
}

func (tp *transactionsPostgres) CreateTransaction(
	ctx context.Context,
	initiatorID int64,
	recipientID int64,
	amount int64,
	description string,
	date time.Time,
	operationTypeID int) (transactionID int64, err error) {

	query, args, err := sq.Insert("transactions").
		Columns("initiator_id", "recipient_id", "amount",
			"description", "date", "operation_type_id").
		Values(initiatorID, recipientID, amount, description, date, operationTypeID).
		Suffix("RETURNING transaction_id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "Insert transaction failed")
	}
	_, err = tp.e.QueryContext(ctx, pg.Scan(&transactionID), query, args...)
	return transactionID, errors.Wrap(err, "Insert transaction failed")
}

func (tp *transactionsPostgres) GetTransactionsByUserID(ctx context.Context, userID int64, limit uint64, timeCursor time.Time) (*[]models.Transaction, error) {
	query, args, err := sq.Select("*").
		From("transactions").
		Where(sq.And{
			sq.Or{
				sq.Eq{"initiator_id": userID}, sq.Eq{"recipient_id": userID}},
			sq.Lt{"date": timeCursor},
		}).OrderBy("date DESC").Limit(limit).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Get transactions by userID failed")
	}
	result := &[]models.Transaction{}
	_, err = tp.e.QueryContext(ctx, result, query, args...)
	return result, errors.Wrap(err, "Get transactions by userID sql failed")
}

func (tp *transactionsPostgres) CountTransactionsLeft(ctx context.Context, userID int64, timeCursor time.Time) (count int, err error) {
	query, args, err := sq.Select("COUNT(*)").
		From("transactions").
		Where(sq.And{
			sq.Or{
				sq.Eq{"initiator_id": userID}, sq.Eq{"recipient_id": userID}},
			sq.Lt{"date": timeCursor},
		}).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "Get transactions count by userID failed")
	}
	_, err = tp.e.QueryContext(ctx, pg.Scan(&count), query, args...)
	return count, errors.Wrap(err, "Get transactions count by userID failed")
}
