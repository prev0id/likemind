-- +goose Up
-- +goose StatementBegin
CREATE TABLE interest (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE user_interest (
    id BIGSERIAL PRIMARY KEY,
    entity_id BIGINT NOT NULL,
    interest_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE group_interest (
    id BIGSERIAL PRIMARY KEY,
    entity_id BIGINT NOT NULL,
    interest_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table interest;
drop table user_interest;
drop table group_interest;
-- +goose StatementEnd
