-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS mb(
    id VARCHAR NOT NULL,
    cust_id VARCHAR NOT NULL,
    acct VARCHAR NOT NULL,
    br VARCHAR NOT NULL,
    segment VARCHAR NOT NULL,
    product VARCHAR NOT NULL,
    cont_code VARCHAR NOT NULL,
    cont_type VARCHAR NOT NULL,
    doc_num VARCHAR,
    subs_loanto VARCHAR,
    line_type VARCHAR,
    end_date VARCHAR,
    amt_tng VARCHAR,
    od_tng VARCHAR,
    stav VARCHAR,
    day_pr_pr VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE mb;
-- +goose StatementEnd
