package models

import "time"

type UserInfo struct {
	UserID  int64 `pg:"user_id,pk"`
	Balance int64 `pg:"balance,notnull"`
}

type Transaction struct {
	TransactionID   int64     `pg:"transaction_id,pk"`
	InitiatorID     int64     `pg:"initiator_id,fk:user_id,notnull,on_update:NO ACTION,on_delete:NO ACTION"`
	RecipientID     int64     `pg:"recipient_id,fk:user_id,notnull,on_update:NO ACTION,on_delete:NO ACTION"`
	Amount          int64     `pg:"amount,notnull"`
	OperationTypeID int       `pg:"operation_type_id"`
	Description     string    `pg:"description"`
	Date            time.Time `pg:"date,notnull"`
}
