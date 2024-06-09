package model

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID            int          `db:"id" json:"id" pk:"true"`
	Name          string       `db:"name" json:"name"`
	Department_ID int          `db:"department_id" json:"department_id"`
	CreatedAt     time.Time    `db:"created_at" json:"created_at"`
	DeletedAt     sql.NullTime `db:"deleted_at" json:"deleted_at,omitempty"`
}
