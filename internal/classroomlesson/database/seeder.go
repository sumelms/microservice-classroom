package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/pkg/seed"
)

func ClassroomLessons() seed.Seed {
	return seed.Seed{
		Name: "CreateClassroomLessons",
		Run: func(db *gorm.DB) error {
			u := &ClassroomLesson{}
			return db.Create(u).Error
		},
	}
}
