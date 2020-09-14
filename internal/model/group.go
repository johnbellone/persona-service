package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string
	Description sql.NullString
}
