package database

import (
	"time"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID `gorm:"type:uuid" sql:"index"`
	ClassroomID uuid.UUID `gorm:"type:uuid" sql:"index"`
	Role        string    `gorm:"size:144"`
}

func (c *Subscription) BeforeCreate(scope *gorm.Scope) error {
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

func (c *Subscription) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		scope.Log("BeforeUpdate error: %v", err)
	}
	return nil
}
