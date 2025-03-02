-- +goose Up
-- +goose StatementBegin
CREATE TABLE interest (
    id bigint PRIMARY KEY NOT NULL,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE TABLE user_interest (
    user_id bigint NOT NULL,
    interest_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    PRIMARY KEY (user_id, interest_id)
);

CREATE TABLE group_interest (
    group_id bigint NOT NULL,
    interest_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    PRIMARY KEY (group_id, interest_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS group_interest;

DROP TABLE IF EXISTS user_interest;

DROP TABLE IF EXISTS interest;
-- +goose StatementEnd
