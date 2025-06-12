-- +goose Up
-- +goose StatementBegin
INSERT INTO users(
    email, password, role, created_at
)VALUES (
        'anon_tech_user@gmail.com', '', 'client', now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email = 'anon_tech_user@gmail.com'
-- +goose StatementEnd
