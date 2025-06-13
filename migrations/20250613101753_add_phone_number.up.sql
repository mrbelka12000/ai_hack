-- +goose Up
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS email;
ALTER TABLE users DROP COLUMN IF EXISTS password;
ALTER TABLE users ADD COLUMN  IF NOT EXISTS phone_number VARCHAR DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
