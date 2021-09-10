package persistent

import "context"

type Storage interface {
	Users() UsersStorage
	Transactions() TransactionsStorage

	RunInTransaction(ctx context.Context, fn func(repository Storage) error) error
}
