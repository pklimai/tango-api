package employee_manager

import "errors"

const (
	FormatErrAgeOutOfRange = "age out of range [%d, %d]"
	FormatMsgAgeOutOfRange = "возраст сотрудника должен находится в диапазоне [%d, %d]"
)

var (
	ErrEmptyName             = errors.New("name should not be empty")
	ErrNotPositiveEmployeeID = errors.New("employee id should be a positive number")
)
