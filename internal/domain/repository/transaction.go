package repository

import "context"

type Transaction interface {
	ExecTx(ctx context.Context, f func(ctx context.Context) error, database string) error
}
