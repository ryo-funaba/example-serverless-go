package mysql

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/enum"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/ryo-funaba/example_echo/internal/utils/logutil"
)

const transactionStructName = "Transaction"

type transaction struct {
	rdb repository.Cluster
}

func NewTransaction(isDebugMode bool) repository.Transaction {
	c := NewCluster(isDebugMode)

	return &transaction{c}
}

func (r *transaction) ExecTx(ctx context.Context, f func(ctx context.Context) error, database string) error {
	logutil.PrintFuncName(transactionStructName, "ExecTx")

	con, err := r.rdb.GetConnection(enum.Primary, database)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	tx, err := con.BeginTx(ctx, nil)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	c := r.rdb.SetTxInCtx(ctx, tx)

	err = f(c)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return err
	}

	return nil
}
