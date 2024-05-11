-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_conf (
    att_conf_id         INTEGER                          PRIMARY KEY,
    att_name            TEXT                            NOT NULL,
    att_conf_type_id    SMALLINT                        NOT NULL,
    att_conf_format_id  SMALLINT                        NOT NULL,
    att_conf_write_id   SMALLINT                        NOT NULL,
    table_name          TEXT                            NOT NULL,
    cs_name             TEXT        DEFAULT ''::TEXT    NOT NULL,
    domain              TEXT        DEFAULT ''::TEXT    NOT NULL,
    family              TEXT        DEFAULT ''::TEXT    NOT NULL,
    member              TEXT        DEFAULT ''::TEXT    NOT NULL,
    name                TEXT        DEFAULT ''::TEXT    NOT NULL,
    ttl                 INTEGER,
    hide                BOOLEAN     DEFAULT false
);

COMMENT ON TABLE att_conf IS 'Attribute Configuration Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS att_conf;
-- +goose StatementEnd
