package persistent

import (
	"balanceManager/models"
	"context"
)

type UsersStorage interface {
	GetUserInfoByID(ctx context.Context, userID int64) (*models.UserInfo, error)
	UpdateBalance(ctx context.Context, userID int64, amount int64) (*models.UserInfo, error)

	CreateUser(ctx context.Context) (*models.UserInfo, error)
}
