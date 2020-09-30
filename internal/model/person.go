package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	Uuid           pgtype.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	GivenName      string
	FamilyName     string
	NickName       string
	AdditionalName string
	BirthDate      pgtype.Date
	BirthPlace     string
	DeathDate      pgtype.Date
	DeathPlace     string
}
