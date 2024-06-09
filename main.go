package main

import (
	"go-x/db_migration"
	"go-x/handlers"
	"go-x/repository"
	"go-x/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbConn, err := db_migration.InitDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	defer dbConn.Close()
	employeeRepo := repository.NewEmployeeRepository(dbConn)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id:[0-9]+}", employeeHandler.GetEmployeeByID).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id:[0-9]+}", employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id:[0-9]+}", employeeHandler.DeleteEmployee).Methods(http.MethodDelete)
	router.HandleFunc("/employees/departments/{id:[0-9]+}", employeeHandler.GetEmployeesByDepartmentID).Methods(http.MethodGet)

	log.Println("server is running at port: 8080")
	http.ListenAndServe(":8080", router)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PONG!"))
}
