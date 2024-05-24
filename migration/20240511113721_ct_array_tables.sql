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

COMMENT ON TABLE att_array_devboolean IS 'Array Boolean Values Table';
-- +goose StatementEnd

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

COMMENT ON TABLE att_array_devdouble IS 'Array Double Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devencoded (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             BYTEA[],
    value_w             BYTEA[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devdouble IS 'Array DevEncoded Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devenum (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r_label       TEXT[],
    value_r             SMALLINT[],
    value_w_label       TEXT[],
    value_w             SMALLINT[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devenum IS 'Array Enum Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devfloat (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             REAL[],
    value_w             REAL[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devfloat IS 'Array Float Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devlong (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             INTEGER[],
    value_w             INTEGER[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devlong IS 'Array Long Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devlong64 (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             BIGINT[],
    value_w             BIGINT[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devlong64 IS 'Array Long64 Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devshort (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             SMALLINT[],
    value_w             SMALLINT[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devshort IS 'Array Short Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devstate (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             INTEGER[],
    value_w             INTEGER[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devstate IS 'Array State Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devstring (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             TEXT[],
    value_w             TEXT[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devstring IS 'Array String Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devuchar (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             uchar[],
    value_w             uchar[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devuchar IS 'Array UChar Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devulong (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             ulong[],
    value_w             ulong[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devulong IS 'Array ULong Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devulong64 (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             ulong64[],
    value_w             ulong64[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devulong64 IS 'Array ULong64 Values Table';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS att_array_devushort (
    att_conf_id         INTEGER NOT NULL,
    data_time           TIMESTAMP WITH TIME ZONE NOT NULL,
    value_r             ushort[],
    value_w             ushort[],
    quality             SMALLINT,
    att_error_desc_id   INTEGER,
    details             JSON
);

COMMENT ON TABLE att_array_devushort IS 'Array UShort Values Table';
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS att_array_devushort;
DROP TABLE IF EXISTS att_array_devulong64;
DROP TABLE IF EXISTS att_array_devulong;
DROP TABLE IF EXISTS att_array_devuchar;
DROP TABLE IF EXISTS att_array_devstring;
DROP TABLE IF EXISTS att_array_devstate;
DROP TABLE IF EXISTS att_array_devshort;
DROP TABLE IF EXISTS att_array_devlong64;
DROP TABLE IF EXISTS att_array_devlong;
DROP TABLE IF EXISTS att_array_devfloat;
DROP TABLE IF EXISTS att_array_devenum;
DROP TABLE IF EXISTS att_array_devencoded;
DROP TABLE IF EXISTS att_array_devdouble;
DROP TABLE IF EXISTS att_array_devboolean;
-- +goose StatementEnd