package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/pkg/seed"
)

func Lessons() seed.Seed {
	return seed.Seed{
		Name: "CreateLessons",
		Run: func(db *gorm.DB) error {
			u := &Lesson{}
			return db.Create(u).Error
		},
	}
}
