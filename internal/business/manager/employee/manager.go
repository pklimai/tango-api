package employee_manager

type Manager struct {
	employeeRepo employeeRepo
}

func New(employeeRepo employeeRepo) *Manager {
	return &Manager{employeeRepo: employeeRepo}
}
