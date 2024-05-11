-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devboolean (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             BOOLEAN,
    value_w             BOOLEAN,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devboolean IS 'Scalar Boolean Values Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS att_scalar_devboolean;
-- +goose StatementEnd