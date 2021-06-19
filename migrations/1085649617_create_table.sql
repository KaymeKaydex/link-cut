-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS url_container
(
    short_url      varchar(5)                  PRIMARY KEY NOT NULL,
    long_url       text                        UNIQUE NOT NULL
    );

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

drop table if exists url_container;

-- +goose StatementEnd
