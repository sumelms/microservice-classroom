package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/pkg/seed"
)

func Classrooms() seed.Seed {
	return seed.Seed{
		Name: "CreateClassrooms",
		Run: func(db *gorm.DB) error {
			u := &Classroom{}
			return db.Create(u).Error
		},
	}
}
