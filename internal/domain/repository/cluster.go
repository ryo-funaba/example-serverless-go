package repository

import (
	"context"
	"database/sql"

	"github.com/ryo-funaba/example_echo/internal/utils/enum"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Cluster interface {
	GetConnection(conn enum.ConnectionType, database string) (*sql.DB, error)
	AddConnection(conn enum.ConnectionType, database, driver, dns string) error
	AddConnectionForTest(conn enum.ConnectionType, database, driver, dns string) error
	GetBoilCtxExecutor(ctx context.Context, conn enum.ConnectionType, database string) (boil.ContextExecutor, error)
	SetTxInCtx(ctx context.Context, tx *sql.Tx) context.Context
}
