package domain

import "database/sql"

type Role struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" gorm:"not null"`
	Description sql.NullString `json:"description" gorm:""`
}
