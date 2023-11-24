-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE INVOICE_STATUS AS ENUM ('PENDING', 'PAID', 'FAILED');

CREATE TABLE invoices (
    id UUID PRIMARY KEY,
    booking_id UUID NOT NULL,
    amount NUMERIC NOT NULL,
    paid_at TIMESTAMPTZ,
    status INVOICE_STATUS DEFAULT 'PENDING',
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoices;
-- +goose StatementEnd
