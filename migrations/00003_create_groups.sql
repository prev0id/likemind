-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups (
    id          BIGSERIAL       PRIMARY KEY,
    name        TEXT            NOT NULL DEFAULT '',
    description TEXT            NOT NULL DEFAULT '',
    author_id   BIGINT          NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE TABLE posts (
    id          BIGSERIAL       PRIMARY KEY,
    group_id    BIGINT          NOT NULL,
    author_id   BIGINT          NOT NULL,
    content     TEXT            NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_posts_group_id ON group_posts (group_id);

CREATE TABLE comments (
    id          BIGSERIAL       PRIMARY KEY NOT NULL,
    post_id     BIGINT          NOT NULL,
    author_id   BIGINT          NOT NULL,
    content     TEXT            NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_comments_post_id ON comments (post_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_comments_post_id;
DROP TABLE IF EXISTS comments;

DROP INDEX IF EXISTS idx_posts_group_id;
DROP TABLE IF EXISTS posts;

DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
