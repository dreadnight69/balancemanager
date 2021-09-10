package logger

import (
	"balanceManager/balancemanager"
	"context"
	"go.uber.org/zap"
)

type Logger struct {
	balancemanager.Service
	logger zap.Logger
}

func New(s balancemanager.Service, l zap.Logger) balancemanager.Service {
	return Logger{
		Service: s,
		logger:  l,
	}
}

func (l Logger) GetUserInfo(ctx context.Context, userID int64) (*balancemanager.UserInfo, error) {
	l.logger.Info("Request GetUserInfo", zap.Int64("userID", userID))
	resp, err := l.Service.GetUserInfo(ctx, userID)
	if err != nil {
		l.logger.Error("Response GetUserInfo", zap.Error(err))
	} else {
		l.logger.Info("Response GetUserInfo", zap.Reflect("response", resp))
	}
	return resp, err
}

func (l Logger) MakeDeposit(ctx context.Context, userID int64, amount string, description string) (*balancemanager.UserInfo, error) {
	l.logger.Info("Request MakeDeposit", zap.Int64("userID", userID), zap.String("amount", amount), zap.String("description", description))
	resp, err := l.Service.MakeDeposit(ctx, userID, amount, description)
	if err != nil {
		l.logger.Error("Response MakeDeposit", zap.Error(err))
	} else {
		l.logger.Info("Response MakeDeposit", zap.Reflect("response", resp))
	}
	return resp, err
}

func (l Logger) WithdrawFunds(ctx context.Context, userID int64, amount string, description string) (*balancemanager.UserInfo, error) {
	l.logger.Info("Request WithdrawFunds", zap.Int64("userID", userID), zap.String("amount", amount), zap.String("description", description))
	resp, err := l.Service.WithdrawFunds(ctx, userID, amount, description)
	if err != nil {
		l.logger.Error("Response WithdrawFunds", zap.Error(err))
	} else {
		l.logger.Info("Response WithdrawFunds", zap.Reflect("response", resp))
	}
	return resp, err
}

func (l Logger) SendFunds(ctx context.Context, senderID int64, recipientID int64, amount string, description string) (*balancemanager.SendFundsResponse, error) {
	l.logger.Info("Request SendFunds", zap.Int64("senderID", senderID), zap.Int64("recipientID", recipientID), zap.String("amount", amount), zap.String("description", description))
	resp, err := l.Service.SendFunds(ctx, senderID, recipientID, amount, description)
	if err != nil {
		l.logger.Error("Response SendFunds", zap.Error(err))
	} else {
		l.logger.Info("Response SendFunds", zap.Reflect("response", resp))
	}
	return resp, err
}

func (l Logger) GetTransactions(ctx context.Context, userID string, limit string, cursor string) (response *[]balancemanager.TransactionResponse, nextCursorUnix int64, err error) {
	l.logger.Info("Request GetTransactions", zap.String("userID", userID), zap.String("limit", limit), zap.String("cursor", cursor))
	resp, c, err := l.Service.GetTransactions(ctx, userID, limit, cursor)
	if err != nil {
		l.logger.Error("Response GetTransactions", zap.Error(err))
	} else {
		l.logger.Info("Response GetTransactions", zap.Reflect("response", resp))
	}
	return resp, c, err
}

func (l Logger) CreateUser(ctx context.Context) (*balancemanager.UserInfo, error) {
	l.logger.Info("Request CreateUser")
	resp, err := l.Service.CreateUser(ctx)
	if err != nil {
		l.logger.Error("Response CreateUser", zap.Error(err))
	} else {
		l.logger.Info("Response CreateUser", zap.Reflect("response", resp))
	}
	return resp, err
}
