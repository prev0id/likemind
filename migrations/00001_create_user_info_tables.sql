-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          BIGSERIAL       PRIMARY KEY,
    nickname    TEXT            NOT NULL UNIQUE,
    name        TEXT            NOT NULL,
    surname     TEXT            NOT NULL,
    about       TEXT            NOT NULL,
    password    BYTEA           NOT NULL,
    email       TEXT            NOT NULL UNIQUE,
    location    TEXT            NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_email ON users (email);

CREATE INDEX idx_user_nickname ON users (nickname);

CREATE TABLE sessions (
    user_id         BIGINT      NOT NULL,
    token           TEXT        NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_sessions_token ON sessions (token);

CREATE TABLE profile_pictures (
    id          TEXT        PRIMARY KEY NOT NULL,
    user_id     BIGINT      NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_profile_picture_user_id ON profile_pictures (user_id);

CREATE TABLE contacts (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    user_id     BIGINT      NOT NULL,
    platform    TEXT        NOT NULL,
    contact     TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_contacts_user_id ON contacts (user_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contacts;

DROP INDEX IF EXISTS idx_contacts_user_id;

DROP TABLE IF EXISTS profile_pictures;

DROP INDEX IF EXISTS idx_profile_picture_user_id;

DROP TABLE IF EXISTS sessions;

DROP INDEX IF EXISTS idx_sessions_token;

DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_user_login;

DROP INDEX IF EXISTS idx_user_nickname;

-- +goose StatementEnd
