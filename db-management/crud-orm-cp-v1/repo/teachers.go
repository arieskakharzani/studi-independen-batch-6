package repo

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	if hasil := t.db.Create(&data); hasil.Error != nil {
		return errors.New("Error INSERT Teacher")
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	rows, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}
	var listTeacher []model.Teacher
	for rows.Next() {
		t.db.ScanRows(rows, &listTeacher)
	}
	return listTeacher, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	if hasil := t.db.Table("teachers").Where("id = ?", 1).Update("name", name); hasil.Error != nil {
		return errors.New("Error UPDATE Teacher")
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if hasil := t.db.Where("id = ?", id).Delete(&teacher); hasil.Error != nil {
		return errors.New("Error DELETE Teacher")
	}
	return nil // TODO: replace this
}
