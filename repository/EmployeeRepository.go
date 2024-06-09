package repository

import (
	"context"
	"database/sql"
	"log"

	"go-x/model"

	_ "github.com/lib/pq"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (er *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (*model.Employee, error) {
	query := "SELECT id, name, department_id FROM employees WHERE id = $1"

	employee := &model.Employee{}
	err := er.db.QueryRowContext(ctx, query, id).Scan(&employee.ID, &employee.Name, &employee.Department_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Employee not found")
			return nil, err
		}
		log.Printf("Error getting employee by ID: %v", err)
		return nil, err
	}

	log.Printf("Retrieved employee: %+v", employee)
	return employee, nil
}

func (er *EmployeeRepository) CreateEmployee(ctx context.Context, employee *model.Employee) error {
	query := "INSERT INTO employees (name, department_id) VALUES ($1, $2)"

	_, err := er.db.ExecContext(ctx, query, employee.Name, employee.Department_ID)
	if err != nil {
		log.Printf("Error creating employee: %v", err)
		return err
	}

	log.Printf("Created employee: %v", employee.Name)
	return nil
}

func (er *EmployeeRepository) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	query := "UPDATE employees SET name = $1, department_id = $2 WHERE id = $3"

	_, err := er.db.ExecContext(ctx, query, employee.Name, employee.Department_ID, employee.ID)
	if err != nil {
		log.Printf("Error updating employee: %v", err)
		return err
	}

	log.Printf("Updated employee: %+v", employee)
	return nil
}

func (er *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	query := "DELETE FROM employees WHERE id = $1"

	_, err := er.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting employee: %v", err)
		return err
	}

	log.Printf("Deleted employee with ID: %d", id)
	return nil
}
func (er *EmployeeRepository) GetEmployeesByDepartmentID(ctx context.Context, departmentID int) ([]*model.Employee, error) {
	query := "SELECT id, name, department_id, created_at, deleted_at FROM employees WHERE department_id = $1"

	rows, err := er.db.QueryContext(ctx, query, departmentID)
	if err != nil {
		log.Printf("Error fetching employees by department ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		employee := &model.Employee{}
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Department_ID, &employee.CreatedAt, &employee.DeletedAt); err != nil {
			log.Printf("Error scanning employee: %v", err)
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %v", err)
		return nil, err
	}

	log.Printf("Retrieved employees: %+v", employees)
	return employees, nil
}
