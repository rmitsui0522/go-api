package service

import (
	"fmt"
	"go-api/pkg/model"
	"go-api/pkg/repository"
)

type DuplicatedRecordException struct {
	fieldName  string
	fieldValue string
}

func (e *DuplicatedRecordException) Error() string {
	return fmt.Sprintf(
		"Duplicated record Error: %s '%s' is already exists.",
		e.fieldName,
		e.fieldValue,
	)
}

type IEmployeeService interface {
	Exists(*model.Employee) error
}

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func (es *EmployeeService) Exists(e *model.Employee) error {
	check := model.Employee{MailAddress: e.MailAddress}

	_, err := es.repo.Find(&check)

	if err == nil {
		return &DuplicatedRecordException{fieldName: "employee.mail_address", fieldValue: e.MailAddress}
	}

	return nil
}
