-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devdouble (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             DOUBLE PRECISION[],
    value_w             DOUBLE PRECISION[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE public.att_array_devdouble IS 'Array Double Values Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS att_array_devdouble;
-- +goose StatementEnd