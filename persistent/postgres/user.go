package postgres

import (
	"balanceManager/models"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type usersPostgres struct {
	baseModelPostgres
}

func (up *usersPostgres) GetUserInfoByID(ctx context.Context, userID int64) (*models.UserInfo, error) {
	query, args, err := sq.Select("*").
		From("users").
		Where(sq.Eq{
			"user_id": userID,
		}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Get user info by userID failed")
	}
	user := &models.UserInfo{}
	_, err = up.e.QueryOneContext(ctx, user, query, args...)
	return user, errors.Wrap(err, "Get user info by userID failed")
}

func (up *usersPostgres) UpdateBalance(ctx context.Context, userID int64, newBalance int64) (*models.UserInfo, error) {
	query, args, err := sq.Update("users").
		Set("balance", newBalance).
		Where(sq.Eq{"user_id": userID}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Update balance failed")
	}
	user := &models.UserInfo{}
	_, err = up.e.QueryContext(ctx, user, query, args...)
	return user, errors.Wrap(err, "Update balance failed")
}

func (up *usersPostgres) CreateUser(ctx context.Context) (*models.UserInfo, error) {
	query, args, err := sq.Insert("users").
		Columns("balance").
		Values(0).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Create user failed")
	}
	user := &models.UserInfo{}
	_, err = up.e.QueryContext(ctx, user, query, args...)
	return user, errors.Wrap(err, "Create user failed")
}
