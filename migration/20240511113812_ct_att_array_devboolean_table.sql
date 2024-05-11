-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devboolean (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             BOOLEAN[],
    value_w             BOOLEAN[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE public.att_array_devboolean IS 'Array Boolean Values Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS att_array_devboolean;
-- +goose StatementEnd