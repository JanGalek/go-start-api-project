package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Abstract struct {
	ID        uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a *Abstract) BeforeCreate(db *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
