package postgres

import (
	"balanceManager/config"
	"balanceManager/persistent"
	"context"
	"github.com/go-pg/pg/v10"
)

type postgres struct {
	db           *pg.DB
	users        *usersPostgres
	transactions *transactionsPostgres
}

func WithConfig(cfg *config.DB) *postgres {
	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Host + ":" + cfg.Port,
		Database: cfg.Name,
	})
	return New(db)
}

func New(db *pg.DB, es ...executor) *postgres {
	var exec executor = db
	if len(es) != 0 {
		exec = es[0]
	}
	return &postgres{
		db:           db,
		users:        &usersPostgres{baseModelPostgres{e: exec}},
		transactions: &transactionsPostgres{baseModelPostgres{e: exec}},
	}
}

type baseModelPostgres struct {
	e executor
}

type executor interface {
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
}

func (p *postgres) Users() persistent.UsersStorage {
	return p.users
}

func (p *postgres) Transactions() persistent.TransactionsStorage {
	return p.transactions
}

func (p *postgres) RunInTransaction(ctx context.Context, fn func(repository persistent.Storage) error) error {
	tx, err := p.db.BeginContext(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = tx.RollbackContext(ctx)
			panic(err)
		}
	}()

	txRepository := New(p.db, tx)

	if err := fn(txRepository); err != nil {
		_ = tx.RollbackContext(ctx)
		return err
	}
	return tx.CommitContext(ctx)
}
