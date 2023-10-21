-- RESTAURANT_ORDER
-- Tabela que contém dados dos pedidos realizados.

-- TABLE

CREATE TABLE IF NOT EXISTS restaurant_order (
    restaurant_order_id serial
        constraint restaurant_order_pk primary key,
    restaurant_order_customer_id int
        constraint restaurant_order_customer_id_fk
            references customer (customer_id),
    restaurant_order_status_id int
        constraint restaurant_order_status_id_fk
            references status(status_id),
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- Indexes

CREATE INDEX IF NOT EXISTS order_customer_id_idx  ON restaurant_order(restaurant_order_customer_id);
CREATE INDEX IF NOT EXISTS order_status_id_idx ON restaurant_order(restaurant_order_status_id);

-- Triggers

CREATE OR REPLACE TRIGGER tr_restaurant_order_timestamps
    BEFORE UPDATE
    ON public.restaurant_order
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();