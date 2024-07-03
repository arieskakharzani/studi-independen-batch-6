package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type LessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) LessonRepo {
	return LessonRepo{db}
}

func (l LessonRepo) Init(data []model.Lesson) error {
	if hasil := l.db.Create(&data); hasil.Error != nil {
		return hasil.Error
	}
	return nil // TODO: replace this
}
