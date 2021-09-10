package balancemanager

import (
	"balanceManager/models"
	operations "balanceManager/pkg/operation_types"
	"github.com/Rhymond/go-money"
)

func transactionsToServiceResponse(transactions *[]models.Transaction) []TransactionResponse {
	results := make([]TransactionResponse, 0, len(*transactions))
	for _, transaction := range *transactions {
		results = append(results, TransactionResponse{
			TransactionID: transaction.TransactionID,
			InitiatorID:   transaction.InitiatorID,
			RecipientID:   transaction.RecipientID,
			Amount:        money.New(transaction.Amount, money.RUB).Display(),
			Description:   transaction.Description,
			Date:          transaction.Date,
			OperationType: operations.NameByType[transaction.OperationTypeID],
		})
	}
	return results
}
