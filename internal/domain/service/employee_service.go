package service

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EmployeeService interface {
	FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error)
	FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{employeeRepository: r}
}

func (s *employeeService) FindOneByEmpNo(ctx context.Context, conn boil.ContextExecutor, empNo int) (*model.Employee, error) {
	return s.employeeRepository.FindOneByEmpNo(ctx, conn, empNo)
}

func (s *employeeService) FindAllByFirstName(ctx context.Context, conn boil.ContextExecutor, firstName string) ([]*model.Employee, error) {
	return s.employeeRepository.FindAllByFirstName(ctx, conn, firstName)
}
