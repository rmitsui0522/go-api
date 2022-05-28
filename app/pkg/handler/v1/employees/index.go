package employees

import (
	"go-api/pkg/repository"
	"go-api/pkg/service"
)

var repo repository.IEmployeeRepository
var esrv service.IEmployeeService

func init() {
	repo = &repository.EmployeeRepository{}
	esrv = &service.EmployeeService{}
}
