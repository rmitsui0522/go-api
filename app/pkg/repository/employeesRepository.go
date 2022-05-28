package repository

import (
	"time"

	"go-api/pkg/model"
)

type IEmployeeRepository interface {
	FindAll() ([]model.Employee, error)
	Find(*model.Employee) (model.Employee, error)
	Create(*model.Employee) error
	Updates(*model.Employee) (model.Employee, error)
	Delete(*model.Employee) error
}

type EmployeeRepository struct{}

func (r *EmployeeRepository) FindAll() ([]model.Employee, error) {
	var employees []model.Employee

	err := db.Find(&employees).Error

	return employees, err
}

func (r *EmployeeRepository) Find(e *model.Employee) (model.Employee, error) {
	var employee model.Employee

	err := db.Where(e).First(&employee).Error

	return employee, err
}

func (r *EmployeeRepository) Create(e *model.Employee) error {
	return db.Create(e).Error
}

func (r *EmployeeRepository) Updates(e *model.Employee) (model.Employee, error) {
	var employee model.Employee
	e.UpdatedAt = time.Now().Round(time.Second)

	err := db.Where("ID = ?", e.ID).Updates(e).First(&employee).Error

	return employee, err
}

func (r *EmployeeRepository) Delete(e *model.Employee) error {
	var employee model.Employee

	err := db.Where(e).Delete(&employee).Error

	return err
}
