package repository

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EmployeeRepository interface {
	FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error)
	FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error)
}
