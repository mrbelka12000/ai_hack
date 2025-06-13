-- +goose Up
-- +goose StatementBegin
ALTER TABLE dialogs ADD COLUMN  IF NOT EXISTS raw_data JSONB DEFAULT '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE dialogs DROP COLUMN  IF  EXISTS raw_data;
-- +goose StatementEnd
