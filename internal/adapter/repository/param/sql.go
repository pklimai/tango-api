package param_repository

const (
	// 	queryGetParamIDAndTableName = `
	// SELECT
	// 	att_conf_id,
	// 	table_name
	// FROM att_conf
	// WHERE
	// 	domain = $1
	//   AND
	// 	name = $2
	//   AND
	// 	member = $3
	// `

	queryFmtGetParamAsScalar = `
SELECT 
	value_r::TEXT,
	value_w::TEXT,
	data_time
FROM %s
WHERE 
	att_conf_id = $1 
  AND 
	data_time BETWEEN $2 AND $3
  AND 
	value_r IS NOT NULL
  AND 
	value_w IS NOT NULL
ORDER BY data_time
`

	queryFmtGetParamAsArray = `
SELECT 
	value_r::TEXT[],
	value_w::TEXT[],
	data_time
FROM %s
WHERE 
	att_conf_id = $1 
  AND 
	data_time BETWEEN $2 AND $3
  AND 
	value_r IS NOT NULL
  AND 
	value_w IS NOT NULL
ORDER BY data_time
`
)
