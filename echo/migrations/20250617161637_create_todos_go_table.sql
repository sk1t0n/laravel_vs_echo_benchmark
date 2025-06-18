-- +goose Up
CREATE TABLE todos_go (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);
-- +goose StatementBegin
select 'up SQL query'
;
-- +goose StatementEnd
-- +goose Down
DROP TABLE todos_go;
-- +goose StatementBegin
select 'down SQL query'
;
-- +goose StatementEnd


