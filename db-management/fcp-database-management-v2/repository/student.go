package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var student []model.Student
	rows, err := s.db.Table("students").Select("*").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s.db.ScanRows(rows, &student)
	}
	return student, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	data := s.db.Create(&student)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	students := &model.Student{}
	data := s.db.Debug().Model(students).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"name":     student.Name,
			"address":  student.Address,
			"class_id": student.ClassId,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	data := s.db.Delete(&model.Student{}, id)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	student := &model.Student{}
	data := s.db.First(&student, id)
	if data.Error != nil {
		return &model.Student{}, data.Error
	}
	return student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var student []model.Student
	var classes []model.Class
	var studentClasses []model.StudentClass

	if err := s.db.Find(&student).Error; err != nil {
		return nil, err
	}
	if len(student) == 0 {
		return &[]model.StudentClass{}, nil
	}
	if err := s.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	err := s.db.Table("students").
		Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
		Joins("JOIN classes ON students.class_id = classes.id").Scan(&studentClasses).Error
	if err != nil {
		return nil, err
	}

	return &studentClasses, nil // TODO: replace this
}
