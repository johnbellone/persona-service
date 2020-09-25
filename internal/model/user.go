package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Uuid               pgtype.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name               string
	Realm              string
	Email              string
	Password           string
	LastPasswordChange pgtype.Timestamptz
	LastLogin          pgtype.Timestamptz
	LastActive         pgtype.Timestamptz
	ActivatedAt        pgtype.Timestamptz
	Attributes         pgtype.JSONB
}
