package mysql

import (
	"context"
	"database/sql"

	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/enum"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type txCtxKey int

const ctxKey txCtxKey = iota

type cluster struct {
	Pools       map[string]*sql.DB
	isDebugMode bool
}

func NewCluster(isDebugMode bool) repository.Cluster {
	pools := map[string]*sql.DB{}
	c := &cluster{pools, isDebugMode}

	return c
}

func (c *cluster) GetConnection(conn enum.ConnectionType, database string) (*sql.DB, error) {
	con, ok := c.Pools[database]
	if !ok {
		return nil, errorutil.Errorf(errorutil.Unknown, "get a connection from pools failed")
	}

	return con, nil
}

func (c *cluster) AddConnection(conn enum.ConnectionType, database, driver, dns string) error {
	con, err := sql.Open(driver, dns)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	err = con.Ping()
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	boil.SetDB(con)
	boil.DebugMode = c.isDebugMode

	c.Pools[database] = con

	return nil
}

func (c *cluster) AddConnectionForTest(conn enum.ConnectionType, database, driver, dns string) error {
	con, err := sql.Open(driver, dns)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	err = con.Ping()
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	boil.SetDB(con)
	boil.DebugMode = c.isDebugMode

	c.Pools[database] = con

	return nil
}

func (c *cluster) GetBoilCtxExecutor(ctx context.Context, conn enum.ConnectionType, database string) (boil.ContextExecutor, error) {
	tx, ok := ctx.Value(ctxKey).(*sql.Tx)
	if !ok {
		con, err := c.GetConnection(conn, database)
		if err != nil {
			return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return con, nil
	}

	return tx, nil
}

func (c *cluster) SetTxInCtx(ctx context.Context, tx *sql.Tx) context.Context {
	v := context.WithValue(ctx, ctxKey, tx)

	return v
}
