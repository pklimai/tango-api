package employee_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	repo "gitlab.com/zigal0-group/nica/tango-api/internal/adapter/repository"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"gitlab.com/zigal0/architect/pkg/business_error"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(ctx context.Context, id int64) (domain.Employee, error) {
	row := r.db.QueryRowContext(ctx, queryGetByID, id)
	if row.Err() != nil {
		return domain.Employee{}, fmt.Errorf(repo.FormatErrQueryRowContext, row.Err())
	}

	var employee Employee

	err := row.Scan(
		&employee.ID,
		&employee.Name,
		&employee.Age,
		&employee.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Employee{}, business_error.New(
				fmt.Errorf(FormatErrNoEmployee, id),
				fmt.Sprintf("не удалось найти сотрудника с id = %d", id),
				business_error.NotFound,
			)
		}

		return domain.Employee{}, fmt.Errorf(repo.FormatErrRowScan, err)
	}

	return toDomain(employee), nil
}

func (r *Repository) Insert(ctx context.Context, employee domain.Employee) (int64, error) {
	row := r.db.QueryRowContext(ctx, queryInsert, employee.Name, employee.Age)
	if row.Err() != nil {
		return 0, fmt.Errorf(repo.FormatErrQueryRowContext, row.Err())
	}

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf(repo.FormatErrRowScan, err)
	}

	return id, nil
}
