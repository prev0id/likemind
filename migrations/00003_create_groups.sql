-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups (
    id          BIGINT          PRIMARY KEY NOT NULL,
    picture_id  TEXT            NOT NULL DEFAULT '',
    name        TEXT            NOT NULL DEFAULT '',
    alias       TEXT            NOT NULL UNIQUE DEFAULT '',
    description TEXT            NOT NULL DEFAULT '',
    author_id   BIGINT          NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_groups_author_id ON groups (author_id);

CREATE TABLE group_posts (
    id          BIGINT          PRIMARY KEY NOT NULL,
    group_id    BIGINT          NOT NULL,
    author_id   BIGINT          NOT NULL,
    title       TEXT            NOT NULL DEFAULT '',
    content     TEXT            NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_group_posts_group_id ON group_posts (group_id);
CREATE INDEX idx_group_posts_author_id ON group_posts (author_id);

CREATE TABLE group_post_comments (
    id          BIGINT          PRIMARY KEY NOT NULL,
    post_id     BIGINT          NOT NULL,
    author_id   BIGINT          NOT NULL,
    content     TEXT            NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_group_post_comments_post_id ON group_post_comments (post_id);
CREATE INDEX idx_group_post_comments_author_id ON group_post_comments (author_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_group_post_comments_author_id;
DROP INDEX IF EXISTS idx_group_post_comments_post_id;
DROP TABLE IF EXISTS group_post_comments;

DROP INDEX IF EXISTS idx_group_posts_author_id;
DROP INDEX IF EXISTS idx_group_posts_group_id;
DROP TABLE IF EXISTS group_posts;

DROP INDEX IF EXISTS idx_groups_author_id;
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
