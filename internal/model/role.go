package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Uuid        pgtype.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Realm       string
	Name        string
	Description pgtype.Text `gorm:"type:text"`
	Metadata    pgtype.JSONB
}
