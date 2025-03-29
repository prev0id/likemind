-- +goose Up
-- +goose StatementBegin
CREATE TABLE interests (
    id          BIGINT          PRIMARY KEY NOT NULL,
    name        TEXT            NOT NULL UNIQUE,
    description TEXT            NOT NULL DEFAULT '',
    group_id    BIGINT          NOT NULL
);

CREATE TABLE interest_groups (
    id          BIGINT          PRIMARY KEY NOT NULL,
    name        TEXT            NOT NULL UNIQUE
);

CREATE TABLE user_interests (
    user_id     BIGINT          NOT NULL,
    interest_id BIGINT          NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, interest_id)
);

CREATE TABLE group_interests (
    group_id    BIGINT          NOT NULL,
    interest_id BIGINT          NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, interest_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS group_interests;

DROP TABLE IF EXISTS user_interests;

DROP TABLE IF EXISTS interests;

DROP TABLE IF EXISTS interest_groups;
-- +goose StatementEnd
