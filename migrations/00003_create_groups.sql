-- +goose Up
-- +goose StatementBegin
CREATE TABLE group (
    id bigint PRIMARY KEY NOT NULL,
    picture_id text NOT NULL,
    name text NOT NULL,
    alias text NOT NULL UNIQUE,
    description text NOT NULL,
    author_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE INDEX idx_group_author_id ON groups (author_id);

-- +goose StatementBegin
-- +goose Down
-- +goose StatementEnd
DROP INDEX IF EXISTS idx_group_author_id;

DROP TABLE IF EXISTS group;

-- +goose StatementEnd
