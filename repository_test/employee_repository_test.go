package repository_test

import (
	"context"
	"database/sql"
	"go-x/model"
	"go-x/repository"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestEmployeeRepository(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://myuser:mypassword@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	er := repository.NewEmployeeRepository(db)

	t.Run("CreateEmployee", func(t *testing.T) {
		ctx := context.Background()
		newEmployee := &model.Employee{
			ID:         1,
			Name:       "John Doe",
			Department: "Engineering",
		}
		err := er.CreateEmployee(ctx, newEmployee)
		if err != nil {
			t.Errorf("Error creating employee: %v", err)
		}
	})

	t.Run("GetEmployeeByID", func(t *testing.T) {
		ctx := context.Background()
		employee, err := er.GetEmployeeByID(ctx, 1)
		if err != nil {
			t.Errorf("Error getting employee by ID: %v", err)
		}
		if employee == nil {
			t.Error("Employee not found")
		}
	})

	t.Run("UpdateEmployee", func(t *testing.T) {
		ctx := context.Background()
		updatedEmployee := &model.Employee{
			ID:         1,
			Name:       "Jane Doe",
			Department: "Marketing",
		}
		err := er.UpdateEmployee(ctx, updatedEmployee)
		if err != nil {
			t.Errorf("Error updating employee: %v", err)
		}
	})

	t.Run("DeleteEmployee", func(t *testing.T) {
		ctx := context.Background()
		err := er.DeleteEmployee(ctx, 1)
		if err != nil {
			t.Errorf("Error deleting employee: %v", err)
		}
	})
}
