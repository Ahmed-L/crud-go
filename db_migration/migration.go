package db_migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"go-x/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	InitializeTables(db)

	return db, nil
}

func InitializeTables(db *sql.DB) error {
	models := []interface{}{
		model.Employee{},
	}
	for _, m := range models {
		if err := createTable(m, db); err != nil {
			return err
		}
	}

	fmt.Println("Tables created successfully!")
	return nil
}

func createTable(model interface{}, db *sql.DB) error {
	tableName := strings.ToLower(reflect.TypeOf(model).Name()) + "s" // Assumes plural naming convention

	var columns []string
	var primaryKey []string
	fields := reflect.TypeOf(model)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)

		columnName := strings.ToLower(field.Tag.Get("db"))
		if columnName == "" {
			columnName = strings.ToLower(field.Name)
		}
		columnType := getColumnType(field.Type)
		columns = append(columns, columnName+" "+columnType)

		if field.Tag.Get("pk") == "true" {
			primaryKey = append(primaryKey, columnName)
		}
	}

	var primaryKeyStmt string
	if len(primaryKey) > 0 {
		primaryKeyStmt = fmt.Sprintf(", PRIMARY KEY (%s)", strings.Join(primaryKey, ", "))
	}

	createTableStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s%s)", tableName, strings.Join(columns, ","), primaryKeyStmt)
	_, err := db.Exec(createTableStmt)
	if err != nil {
		log.Fatalf("Error creating table %s: %v", tableName, err)
	}

	fmt.Printf("Table %s created successfully\n", tableName)
	return nil
}

func DeleteTables(db *sql.DB) {
	for _, model := range []interface{}{model.Employee{}} {
		dropTable(model, db)
	}

	fmt.Println("All tables dropped successfully")
}

func dropTable(model interface{}, db *sql.DB) {
	tableName := strings.ToLower(reflect.TypeOf(model).Name()) + "s" // Assumes plural naming convention

	dropTableStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)

	_, err := db.Exec(dropTableStmt)
	if err != nil {
		log.Fatalf("Error dropping table %s: %v", tableName, err)
	}

	fmt.Printf("Table %s dropped successfully\n", tableName)
}

func getColumnType(fieldType reflect.Type) string {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.String:
		return "VARCHAR(255)"
	default:
		return "TEXT"
	}
}
