-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devlong (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             INTEGER,
    value_w             INTEGER,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devlong IS 'Scalar Long Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devlong64 (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             BIGINT,
    value_w             BIGINT,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devlong64 IS 'Scalar Long64 Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devshort (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             smallint,
    value_w             smallint,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devshort IS 'Scalar Short Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devstring (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             TEXT,
    value_w             TEXT,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devstring IS 'Scalar String Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devstate (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             INTEGER,
    value_w             INTEGER,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devstate IS 'Scalar State Values Table';
-- +goose StatementEnd

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

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devencoded (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             BYTEA,
    value_w             BYTEA,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devencoded IS 'Scalar DevEncoded Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devenum (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r_label       TEXT,
    value_r             SMALLINT,
    value_w_label       TEXT,
    value_w             SMALLINT,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devenum IS 'Scalar Enum Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devfloat (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             REAL,
    value_w             REAL,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devfloat IS 'Scalar Float Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devuchar (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             uchar,
    value_w             uchar,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devuchar IS 'Scalar UChar Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devulong (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             ulong,
    value_w             ulong,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devulong IS 'Scalar ULong Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devulong64 (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             ulong64,
    value_w             ulong64,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devulong64 IS 'Scalar ULong64 Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_scalar_devushort (
    att_conf_id         INTEGER                     NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE    NOT NULL,
    value_r             ushort,
    value_w             ushort,
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_scalar_devushort IS 'Scalar UShort Values Table';
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS att_scalar_devushort;
DROP TABLE IF EXISTS att_scalar_devulong64;
DROP TABLE IF EXISTS att_scalar_devulong;
DROP TABLE IF EXISTS att_scalar_devuchar;
DROP TABLE IF EXISTS att_scalar_devfloat;
DROP TABLE IF EXISTS att_scalar_devenum;
DROP TABLE IF EXISTS att_scalar_devencoded;
DROP TABLE IF EXISTS att_scalar_devdouble;
DROP TABLE IF EXISTS att_scalar_devboolean;
DROP TABLE IF EXISTS att_scalar_devstate;
DROP TABLE IF EXISTS att_scalar_devstring;
DROP TABLE IF EXISTS att_scalar_devshort;
DROP TABLE IF EXISTS att_scalar_devlong64;
DROP TABLE IF EXISTS att_scalar_devlong;
-- +goose StatementEnd

