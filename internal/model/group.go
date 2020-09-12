package model

import (
	"github.com/go-gorm/gorm"
)

type Group struct {
	gorm.Model
	Name        string
	Description sql.NullString
}
