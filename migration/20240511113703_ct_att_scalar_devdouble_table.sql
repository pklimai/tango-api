-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devdouble (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             DOUBLE PRECISION,
    value_w             DOUBLE PRECISION,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devdouble IS 'Scalar Double Values Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS att_scalar_devdouble;
-- +goose StatementEnd