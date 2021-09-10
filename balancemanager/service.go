package balancemanager

import (
	"balanceManager/models"
	"balanceManager/persistent"
	operations "balanceManager/pkg/operation_types"
	"context"
	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
	"sort"
	"strconv"
	"time"
)

type service struct {
	persistent persistent.Storage
}

func New(p persistent.Storage) Service {
	return service{
		persistent: p,
	}
}

func (s service) GetUserInfo(ctx context.Context, userID int64) (*UserInfo, error) {
	user, err := s.persistent.Users().GetUserInfoByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		UserID:  user.UserID,
		Balance: money.New(user.Balance, money.RUB).Display(),
	}, nil
}

func (s service) MakeDeposit(ctx context.Context, userID int64, amount string, description string) (*UserInfo, error) {
	var updatedUser *models.UserInfo
	depositAmount, err := monetaryValueFromString(amount)
	if err != nil {
		return nil, err
	}
	err = s.persistent.RunInTransaction(ctx, func(repo persistent.Storage) (err error) {
		user, err := repo.Users().GetUserInfoByID(ctx, userID)
		if err != nil {
			return err
		}
		balance := user.Balance
		updatedUser, err = repo.Users().UpdateBalance(ctx, userID, balance+depositAmount)
		if err != nil {
			return err
		}
		_, err = s.persistent.Transactions().CreateTransaction(ctx, userID, userID, depositAmount, description, time.Now(), operations.DEPOSIT)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		UserID:  updatedUser.UserID,
		Balance: money.New(updatedUser.Balance, money.RUB).Display(),
	}, nil
}

func (s service) WithdrawFunds(ctx context.Context, userID int64, amount string, description string) (*UserInfo, error) {
	var updatedUser *models.UserInfo
	depositAmount, err := monetaryValueFromString(amount)
	if err != nil {
		return nil, err
	}
	err = s.persistent.RunInTransaction(ctx, func(repo persistent.Storage) (err error) {
		user, err := repo.Users().GetUserInfoByID(ctx, userID)
		if err != nil {
			return err
		}
		balance := user.Balance
		if balance-depositAmount < 0 {
			return errors.New("Not enough funds")
		}
		updatedUser, err = repo.Users().UpdateBalance(ctx, userID, balance-depositAmount)
		if err != nil {
			return err
		}
		_, err = s.persistent.Transactions().CreateTransaction(ctx, userID, userID, depositAmount, description, time.Now(), operations.WITHDRAW)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		UserID:  updatedUser.UserID,
		Balance: money.New(updatedUser.Balance, money.RUB).Display(),
	}, nil
}

func (s service) SendFunds(ctx context.Context, senderID int64, recipientID int64, amount string, description string) (*SendFundsResponse, error) {
	var updatedSenderInfo, updatedRecipientInfo *models.UserInfo
	depositAmount, err := monetaryValueFromString(amount)
	if err != nil {
		return nil, err
	}
	err = s.persistent.RunInTransaction(ctx, func(repo persistent.Storage) (err error) {
		senderInfo, err := repo.Users().GetUserInfoByID(ctx, senderID)
		if err != nil {
			return err
		}
		if senderInfo.Balance-depositAmount < 0 {
			return errors.New("Not enough funds")
		}
		recipientInfo, err := repo.Users().GetUserInfoByID(ctx, recipientID)
		if err != nil {
			return err
		}
		updatedSenderInfo, err = repo.Users().UpdateBalance(ctx, senderID, senderInfo.Balance-depositAmount)
		if err != nil {
			return err
		}
		updatedRecipientInfo, err = repo.Users().UpdateBalance(ctx, recipientID, recipientInfo.Balance+depositAmount)
		if err != nil {
			return err
		}
		_, err = s.persistent.Transactions().CreateTransaction(ctx, senderID, recipientID, depositAmount, description, time.Now(), operations.TRANSFER)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &SendFundsResponse{
		SenderID:      updatedSenderInfo.UserID,
		SenderBalance: money.New(updatedSenderInfo.Balance, money.RUB).Display(),
		RecipientID:   updatedRecipientInfo.UserID,
		Amount:        amount,
	}, nil
}

func (s service) GetTransactions(ctx context.Context, userID string, limit string, cursor string) (response *[]TransactionResponse, nextCursorUnix int64, err error) {
	timestamp, err := strconv.ParseInt(cursor, 10, 64)
	if err != nil {
		timestamp = time.Now().Unix()
	}
	userIDint, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	timeCursor := time.Unix(timestamp, 0)
	uintLimit, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return nil, 0, err
	}

	transactions, err := s.persistent.Transactions().GetTransactionsByUserID(ctx, userIDint, uintLimit, timeCursor)
	if err != nil {
		return nil, 0, err
	}
	if len(*transactions) == 0 {
		return nil, 0, errors.New("user doesn't exist")
	}

	results := transactionsToServiceResponse(transactions)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Date.After(results[j].Date)
	})

	nextCursor := results[len(results)-1].Date
	countTransactionsLeft, err := s.persistent.Transactions().CountTransactionsLeft(ctx, userIDint, nextCursor)
	if err != nil {
		return nil, 0, err
	}

	nextCursorUnix = nextCursor.Unix()
	if countTransactionsLeft == 0 {
		nextCursorUnix = 0
	}

	return &results, nextCursorUnix, nil
}

func (s service) CreateUser(ctx context.Context) (*UserInfo, error) {
	user, err := s.persistent.Users().CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		UserID:  user.UserID,
		Balance: money.New(user.Balance, money.RUB).Display(),
	}, nil
}
