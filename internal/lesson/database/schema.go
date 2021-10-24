package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// lesson struct
type Lesson struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title       string    `gorm:"size:100"`
	Subtitle    string    `gorm:"size:100"`
	Excerpt     string    `gorm:"size:144"`
	Description string    `gorm:"size:255"`
	Module      string	  `gorm:"size:155"`
	SubjectID   uuid.UUID `gorm:"type:uuid" sql:"index"`
}

func (c *Lesson) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	scope.SetColumn("UUID", id.String()) // nolint: errcheck

	if c.UpdatedAt.IsZero() {
		err = scope.SetColumn("UpdatedAt", time.Now())
		if err != nil {
			scope.Log("BeforeCreate error: %v", err)
		}
	}

	err = scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeCreate error: %v", err)
	}
	return nil
}

func (c *Lesson) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeUpdate error: %v", err)
	}
	return nil
}
