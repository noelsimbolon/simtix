-- +goose Up
-- +goose StatementBegin
ALTER TABLE seats ADD COLUMN booking_id UUID NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE seats DROP COLUMN booking_id;
-- +goose StatementEnd
