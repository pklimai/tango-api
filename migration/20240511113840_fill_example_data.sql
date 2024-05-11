-- +goose Up
-- +goose StatementBegin
INSERT INTO att_conf (att_conf_id, att_name, att_conf_type_id, att_conf_format_id, att_conf_write_id, table_name, cs_name, domain, family, member, name, ttl, hide) VALUES
    (1, 'att_name_sd', 5, 1, 1, 'att_scalar_devdouble', 'cs_name_sd', 'domain_sd', 'family_sd', 'member_sd', 'name_sd', 0, false),
    (2, 'att_name_ad', 15, 1, 1, 'att_array_devdouble', 'cs_name_ad', 'domain_ad', 'family_ad', 'member_ad', 'name_ad', 0, false),
    (3, 'att_name_sb', 1, 1, 1, 'att_scalar_devboolean', 'cs_name_sb', 'domain_sb', 'family_sb', 'member_sb', 'name_sb', 0, false),
    (4, 'att_name_ss', 8, 1, 1, 'att_scalar_devstring', 'cs_name_ss', 'domain_ss', 'family_ss', 'member_ss', 'name_ss', 0, false),
    (5, 'att_name_ss', 8, 1, 1, 'att_scalar_devstring', 'cs_name_ss', 'domain_ss', 'family_ss', 'member_ss', 'name_ss', 0, false),
    (6, 'att_name_si32', 3, 1, 1, 'att_scalar_devshort', 'cs_name_si32', 'domain_si32', 'family_si32', 'member_si32', 'name_si32', 0, false);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO att_scalar_devdouble (att_conf_id, data_time, value_r, value_w) VALUES
    (1, NOW(), 13.37, 13.37),
    (1, NOW(), 14.37, 14.37),
    (2, NOW(), 13.37, 13.37);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO att_array_devdouble (att_conf_id, data_time, value_r, value_w) VALUES
    (2, NOW(), ARRAY[13.37, 14.37]::DOUBLE PRECISION[], ARRAY[13.37, 14.37]::DOUBLE PRECISION[]),
    (2, NOW(), ARRAY[14.37, 13.37]::DOUBLE PRECISION[], ARRAY[14.37, 13.37]::DOUBLE PRECISION[]),
    (1, NOW(), ARRAY[13.37]::DOUBLE PRECISION[], ARRAY[13.37]::DOUBLE PRECISION[]);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO att_scalar_devboolean (att_conf_id, data_time, value_r, value_w) VALUES
    (3, NOW(), true, false),
    (3, NOW(), false, true),
    (1, NOW(), false, false);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO att_array_devboolean (att_conf_id, data_time, value_r, value_w) VALUES
    (2, NOW(), ARRAY[true, false]::BOOLEAN[], ARRAY[true, false]::BOOLEAN[]),
    (2, NOW(), ARRAY[false, true]::BOOLEAN[], ARRAY[false, true]::BOOLEAN[]),
    (1, NOW(), ARRAY[false, false]::BOOLEAN[], ARRAY[false, false]::BOOLEAN[]);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE att_array_devboolean
TRUNCATE TABLE att_scalar_devboolean
TRUNCATE TABLE att_array_devdouble
TRUNCATE TABLE att_scalar_devdouble
TRUNCATE TABLE att_conf
-- +goose StatementEnd
