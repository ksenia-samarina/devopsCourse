-- +goose Up
-- +goose StatementBegin
alter table orders add column if not exists created_at timestamp with time zone default now() not null;
CREATE INDEX IF NOT EXISTS orders_status_created_at ON orders (status, created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders_status_created_at;
alter table orders drop column if exists created_at;
-- +goose StatementEnd
