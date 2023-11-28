-- +goose Up
-- +goose StatementBegin
ALTER TABLE seats
    ADD COLUMN seat_row VARCHAR(2) NOT NULL;
ALTER TABLE seats
    ADD COLUMN seat_number NUMERIC NOT NULL;
ALTER TABLE seats
    ADD COLUMN price NUMERIC DEFAULT 0;
ALTER TABLE events
    ADD COLUMN event_time TIMESTAMPTZ NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Down statement for seats table
ALTER TABLE events
    DROP COLUMN event_date;

ALTER TABLE seats
    DROP COLUMN price;
ALTER TABLE seats
    DROP COLUMN seat_row;
ALTER TABLE seats
    DROP COLUMN seat_number;

-- +goose StatementEnd
