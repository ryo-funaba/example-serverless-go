package usecase

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/service"
	"github.com/ryo-funaba/example_echo/internal/usecase/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EmployeeUsecase interface {
	FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error)
	FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error)
}

type employeeUsecase struct {
	employeeService service.EmployeeService
}

func NewEmployeeUsecase(s service.EmployeeService) EmployeeUsecase {
	return &employeeUsecase{employeeService: s}
}

func (u *employeeUsecase) FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error) {
	e, err := u.employeeService.FindOneByEmpNo(ctx, conn, empNo)
	if err != nil {
		return nil, err
	}

	return model.EmployeeFromDomainModel(e), nil
}

func (u *employeeUsecase) FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error) {
	e, err := u.employeeService.FindAllByFirstName(ctx, conn, firstName)
	if err != nil {
		return nil, err
	}

	s := make([]*model.Employee, 0, len(e))

	for _, v := range e {
		s = append(s, model.EmployeeFromDomainModel(v))
	}

	return s, nil
}
