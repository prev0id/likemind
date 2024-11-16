-- +goose Up
-- +goose StatementBegin
CREATE TABLE interest (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE user_interest (
    user_id     BIGINT NOT NULL,
    interest_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, interest_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table interest;
drop table user_interest;
-- +goose StatementEnd
