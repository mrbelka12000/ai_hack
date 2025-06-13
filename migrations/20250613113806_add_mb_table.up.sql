-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS personal_data(
    id VARCHAR NOT NULL,
    br VARCHAR NOT NULL,
    currency VARCHAR,
    beg_date VARCHAR,
    env_date VARCHAR,
    prol_date VARCHAR,
    prol_count VARCHAR,
    amt VARCHAR,
    amt_tng VARCHAR,
    od VARCHAR,
    pr_od VARCHAR,
    day_pr_od VARCHAR,
    pgg VARCHAR,
    stav VARCHAR,
    sht VARCHAR,
    br_vyd VARCHAR,
    flwork VARCHAR,
    rate_effective VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE mb;
-- +goose StatementEnd
