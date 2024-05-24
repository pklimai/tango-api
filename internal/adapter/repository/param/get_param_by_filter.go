package param_repository

import (
	"context"
	"fmt"
	"slices"
	"time"

	sq "github.com/Masterminds/squirrel"
	repo "gitlab.com/zigal0-group/nica/tango-api/internal/adapter/repository"
	tool_collection "gitlab.com/zigal0-group/nica/tango-api/internal/business/tool/collection"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"gitlab.com/zigal0/architect/pkg/business_error"
)

func (r *Repository) GetParamByFilter(
	ctx context.Context,
	filter domain.ParamFilter,
) (domain.Params, error) {
	conditions := sq.And{}

	conditions = append(conditions, sq.Eq{"domain": filter.Domain})
	conditions = append(conditions, sq.Eq{"name": filter.Name})

	if filter.Member != "" {
		conditions = append(conditions, sq.Eq{"member": filter.Member})
	}

	qb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("att_conf_id", "att_conf_type_id", "table_name").
		From("att_conf").
		Where(conditions)

	sql, args, err := qb.ToSql()
	if err != nil {
		return domain.Params{}, fmt.Errorf(repo.FmtErrToSQL, err)
	}

	var searchOptionsSlice []SearchOptions

	err = r.db.SelectContext(ctx, &searchOptionsSlice, sql, args...)
	if err != nil {
		return domain.Params{}, fmt.Errorf(repo.FmtErrSelectContext, err)
	}

	if len(searchOptionsSlice) == 0 {
		return domain.Params{}, business_error.New(
			fmt.Errorf(FmtErrNoTable, filter),
			fmt.Sprintf("не удалось найти таблицу по парметру для фильтра: %v", filter),
			business_error.NotFound,
		)
	}

	if len(searchOptionsSlice) > 1 {
		return domain.Params{}, business_error.New(
			fmt.Errorf(FmtErr2MoreTables, filter),
			fmt.Sprintf("нашлось 2 или более таблиц для фильтра: %v", filter),
			business_error.FailedPrecondition,
		)
	}

	tableName := searchOptionsSlice[0].TableName
	attConfTypeID := searchOptionsSlice[0].AttConfTypeID
	attConfID := searchOptionsSlice[0].AttConfID

	if !slices.Contains(allTables, tableName) {
		return domain.Params{}, business_error.New(
			fmt.Errorf(FmtErrUnimplimentedForTable, tableName),
			fmt.Sprintf("функционал для таблицы = '%s' отсутствует", tableName),
			business_error.Unimplemented,
		)
	}

	res := domain.Params{
		Type: paramTypeToKind(attConfTypeID),
	}

	if slices.Contains(scalarTables, tableName) {
		res.ScalarParams, err = r.getScalar(ctx, tableName, attConfID, filter.TimeFrom, filter.TimeTo)
		if err != nil {
			return domain.Params{}, fmt.Errorf("getScalar: %w", err)
		}
	}

	if slices.Contains(arrayTables, tableName) {
		res.ArrayParams, err = r.getArray(ctx, tableName, attConfID, filter.TimeFrom, filter.TimeTo)
		if err != nil {
			return domain.Params{}, fmt.Errorf("getArray: %w", err)
		}
	}

	return res, nil
}

func (r *Repository) getScalar(
	ctx context.Context,
	tableName string,
	attConfID int32,
	timeFrom time.Time,
	timeTo time.Time,
) ([]domain.ScalarParam, error) {
	var params []ScalarParam

	if timeTo.IsZero() {
		timeTo = time.Now()
	}

	err := r.db.SelectContext(
		ctx,
		&params,
		fmt.Sprintf(queryFmtGetParamAsScalar, tableName), // TODO: sql injection
		attConfID,
		timeFrom,
		timeTo,
	)
	if err != nil {
		return nil, fmt.Errorf(repo.FmtErrSelectContext, err)
	}

	return tool_collection.Map(params, scalarParamToDomain), nil
}

func (r *Repository) getArray(
	ctx context.Context,
	tableName string,
	attConfID int32,
	timeFrom time.Time,
	timeTo time.Time,
) ([]domain.ArrayParam, error) {
	var params []ArrayParam

	if timeTo.IsZero() {
		timeTo = time.Now()
	}

	err := r.db.SelectContext(
		ctx,
		&params,
		fmt.Sprintf(queryFmtGetParamAsArray, tableName), // TODO: sql injection
		attConfID,
		timeFrom,
		timeTo,
	)
	if err != nil {
		return nil, fmt.Errorf(repo.FmtErrSelectContext, err)
	}

	return tool_collection.Map(params, arrayParamToDomain), nil
}
