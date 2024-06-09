package service

import (
	"context"
	"go-x/model"
	"go-x/repository"
)

type EmployeeService interface {
	GetEmployeeByID(ctx context.Context, id int) (*model.Employee, error)
	CreateEmployee(ctx context.Context, employee *model.Employee) error
	UpdateEmployee(ctx context.Context, employee *model.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesByDepartmentID(ctx context.Context, department_id int) ([]*model.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: *repo}
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id int) (*model.Employee, error) {
	return s.repo.GetEmployeeByID(ctx, id)
}

func (s *employeeService) CreateEmployee(ctx context.Context, employee *model.Employee) error {
	return s.repo.CreateEmployee(ctx, employee)
}

func (s *employeeService) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	return s.repo.UpdateEmployee(ctx, employee)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (s *employeeService) GetEmployeesByDepartmentID(ctx context.Context, department_id int) ([]*model.Employee, error) {
	return s.repo.GetEmployeesByDepartmentID(ctx, department_id)
}
