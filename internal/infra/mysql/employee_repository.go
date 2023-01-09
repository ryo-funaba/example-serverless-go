package mysql

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type employeeRepository struct{}

func NewEmployeeRepository() repository.EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error) {
	q := []qm.QueryMod{
		model.EmployeeWhere.EmpNo.EQ(empNo),
	}

	row, err := model.Employees(q...).One(ctx, conn)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return row, nil
}

func (r *employeeRepository) FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error) {
	q := []qm.QueryMod{
		model.EmployeeWhere.FirstName.EQ(firstName),
	}

	rows, err := model.Employees(q...).All(ctx, conn)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return rows, nil
}
