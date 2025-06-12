-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS dialogs(
    id UUID NOT NULL PRIMARY KEY,
    client_id INTEGER NOT NULL,
    operator_id INTEGER,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT  now(),
    FOREIGN KEY (client_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS dialogs_messages(
    dialog_id UUID NOT NULL,
    message VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT  now(),
    FOREIGN KEY (dialog_id) REFERENCES dialogs(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dialogs_messages;
DROP TABLE IF EXISTS dialogs;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
