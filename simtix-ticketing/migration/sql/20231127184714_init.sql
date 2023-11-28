-- +goose Up
-- +goose StatementBegin
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE events
(
    id         UUID PRIMARY KEY,
    event_name VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TYPE SEAT_STATUS AS ENUM ('OPEN', 'ONGOING', 'BOOKED');

CREATE TABLE seats
(
    id         UUID PRIMARY KEY,
    event_id   UUID,
    status     SEAT_STATUS DEFAULT 'OPEN',
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY(event_id) REFERENCES events(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE seats;
DROP TABLE events;
-- +goose StatementEnd
