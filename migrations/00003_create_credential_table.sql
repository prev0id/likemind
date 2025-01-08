-- +goose Up
-- +goose StatementBegin
CREATE TABLE credential (
    id BIGSERIAL PRIMARY KEY,
    user_id TEXT DEFAULT '' NOT NULL,
    login TEXT DEFAULT '' NOT NULL,
    password TEXT DEFAULT '' NOT NULL,
    uuid TEXT DEFAULT '' NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table credential;
-- +goose StatementEnd
