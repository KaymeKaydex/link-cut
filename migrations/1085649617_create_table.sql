-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS urls
(
    short      varchar(5)                   PRIMARY KEY NOT NULL,
    long       varchar(100)                 PRIMARY KEY NOT NULL,
    used_at    timestamp with time zone     NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_short_urls
    ON urls (short);

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

drop table if exists news;
drop index if exists idx_news_deleted_at;

-- +goose StatementEnd
