-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_subscriptions (
    user_id         BIGINT      NOT NULL,
    group_id        BIGINT      NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_subscriptions ON user_subscriptions (user_id, group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX ID EXISTS idx_group_id_user_id;

DROP TABLE IF EXISTS user_subscriptions;
-- +goose StatementEnd
