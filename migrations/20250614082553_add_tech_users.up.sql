-- +goose Up
-- +goose StatementBegin
    INSERT INTO users(
role, created_at, phone_number
    ) VALUES (
              'client', now(), '8888'
);

INSERT INTO users(
    role, created_at, phone_number
) VALUES (
'operator', now(), '7777'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users where phone_number = '8888';
DELETE FROM users where phone_number = '7777';
-- +goose StatementEnd
