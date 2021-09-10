package http

import (
	"balanceManager/balancemanager"
)

func userInfoToHttpResponse(info *balancemanager.UserInfo) UserInfoResponse {
	return UserInfoResponse{
		UserID:  info.UserID,
		Balance: info.Balance,
	}
}

func sendFundsResponseToHttp(response *balancemanager.SendFundsResponse) SendFundsResponse {
	return SendFundsResponse{
		SenderID:      response.SenderID,
		SenderBalance: response.SenderBalance,
		RecipientID:   response.RecipientID,
		Amount:        response.Amount,
	}
}

func transactionsToHttpResponse(transactions *[]balancemanager.TransactionResponse, nextCursor int64) GetTransactionsResponse {
	result := make([]Transaction, 0, len(*transactions))
	for _, transaction := range *transactions {
		result = append(result, Transaction{
			TransactionID: transaction.TransactionID,
			InitiatorID:   transaction.InitiatorID,
			RecipientID:   transaction.RecipientID,
			Amount:        transaction.Amount,
			Description:   transaction.Description,
			Date:          transaction.Date,
			OperationType: transaction.OperationType,
		})
	}
	return GetTransactionsResponse{
		Transactions: result,
		NextCursor:   nextCursor,
	}
}
