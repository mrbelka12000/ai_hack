-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS personal_data(
    call_id VARCHAR DEFAULT '',
    phone_number VARCHAR DEFAULT '',
    br VARCHAR DEFAULT '',
    currency VARCHAR DEFAULT '',
    beg_date VARCHAR DEFAULT '',
    end_date VARCHAR DEFAULT '',
    prol_date VARCHAR DEFAULT '',
    prol_count VARCHAR DEFAULT '',
    amt VARCHAR DEFAULT '',
    amt_tng VARCHAR DEFAULT '',
    od VARCHAR DEFAULT '',
    pr_od VARCHAR DEFAULT '',
    day_pr_od VARCHAR DEFAULT '',
    pog VARCHAR DEFAULT '',
    stav VARCHAR DEFAULT '',
    sht VARCHAR DEFAULT '',
    br_vyd VARCHAR DEFAULT '',
    flwork VARCHAR DEFAULT '',
    rate_effective VARCHAR DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE personal_data;
-- +goose StatementEnd
