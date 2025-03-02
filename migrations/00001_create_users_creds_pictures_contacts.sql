-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id bigint PRIMARY KEY NOT NULL,
    nickname text NOT NULL UNIQUE,
    name text NOT NULL,
    surname text NOT NULL,
    about text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE TABLE credentials (
    id text PRIMARY KEY NOT NULL,
    password bytea NOT NULL,
    login text NOT NULL UNIQUE,
    user_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE INDEX idx_credentials_user_id ON credentials (user_id);

CREATE TABLE profile_picture (
    id text PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE INDEX idx_profile_picture_user_id ON profile_picture (user_id);

CREATE TABLE contact (
    id bigint PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    platform text NOT NULL,
    contact text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE INDEX idx_contacts_user_id ON contact (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contact;

DROP INDEX IF EXISTS idx_contact_user_id;

DROP TABLE IF EXISTS profile_picture;

DROP INDEX IF EXISTS idx_profile_picture_user_id;

DROP TABLE IF EXISTS credential;

DROP INDEX IF EXISTS idx_credential_user_id;

DROP TABLE IF EXISTS user;
-- +goose StatementEnd
