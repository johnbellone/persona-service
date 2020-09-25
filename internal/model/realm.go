package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Realm struct {
	gorm.Model

	Uuid        pgtype.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string
	Description pgtype.Text `gorm:"type:text"`
	Attributes  pgtype.JSONB
}
