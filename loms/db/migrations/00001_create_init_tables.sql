-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks
(
    sku          bigint  NOT NULL,
    warehouse_id bigint  NOT NULL,
    count        integer NOT NULL,
    CONSTRAINT stock_pk PRIMARY KEY (sku, warehouse_id)
);

-- TODO analyze additional indexes

CREATE TABLE IF NOT EXISTS orders
(
    order_id bigserial PRIMARY KEY,
    user_id  bigint   NOT NULL,
    status   smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS stocks_reservation
(
    order_id     bigint   NOT NULL,
    sku          bigint   NOT NULL,
    warehouse_id bigint   NOT NULL,
    count        integer  NOT NULL,
    status       smallint NOT NULL,
    CONSTRAINT stocks_reservation_pk PRIMARY KEY (order_id, sku, warehouse_id)
);

CREATE INDEX IF NOT EXISTS stocks_reservation_sku ON stocks_reservation (sku);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX stocks_reservation_sku;
DROP TABLE IF EXISTS stocks_reservation;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
