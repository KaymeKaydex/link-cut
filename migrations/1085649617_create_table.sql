-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS urls
(
    short_url      varchar(5)                  PRIMARY KEY NOT NULL,
    long_url       varchar(100)                UNIQUE NOT NULL
    );

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

drop table if exists urls;
drop index if exists idx_urls;

-- +goose StatementEnd
