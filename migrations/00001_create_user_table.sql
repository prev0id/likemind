-- +goose Up
-- +goose StatementBegin
create table "user"(
    id       BIGSERIAL PRIMARY KEY,
    nickname TEXT DEFAULT '' NOT NULL,
    name     TEXT DEFAULT '' NOT NULL,
    surname  TEXT DEFAULT '' NOT NULL,
    pfp_id   TEXT DEFAULT '' NOT NULL,
    about    TEXT DEFAULT '' NOT NULL,
    contacts TEXT[] DEFAULT ARRAY[]::TEXT[] NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user";
-- +goose StatementEnd
