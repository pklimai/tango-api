package employee_repository

const (
	queryGetByID = `
SELECT 
	id, 
	name, 
	age,
	created_at 
FROM employee 
WHERE id = $1
`

	queryInsert = `
INSERT INTO employee (name, age) VALUES ($1, $2)
RETURNING id
`
)
