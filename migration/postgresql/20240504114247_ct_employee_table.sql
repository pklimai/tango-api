-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee (
    id          BIGSERIAL   PRIMARY KEY,
    name        TEXT        NOT NULL,
    age         INTEGER     NOT NULL,
    created_at  TIMESTAMP   WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd
