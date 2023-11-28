-- +goose Up
-- +goose StatementBegin
INSERT INTO events (id, event_name, created_at, updated_at, deleted_at)
VALUES
    ('d3407173-3984-460e-8429-327b878667ff', 'Event 1', NOW(), NOW(), NULL),
    ('d1ef8e44-dd8d-4002-9c5a-0520b3fefcfd', 'Event 2', NOW(), NOW(), NULL),
    ('8ce58cdc-3f20-419f-8b61-c3ecf4aa7975', 'Event 3', NOW(), NOW(), NULL);

INSERT INTO seats (id, event_id, status, created_at, updated_at, deleted_at)
VALUES
    ('0a526788-724b-4bf6-9521-56b2afa2a584', 'd3407173-3984-460e-8429-327b878667ff', 'OPEN', NOW(), NOW(), NULL),
    ('74fffff4-4f46-46e0-8099-14486373296a', 'd3407173-3984-460e-8429-327b878667ff', 'OPEN', NOW(), NOW(), NULL),
    ('05ded810-dc4b-461b-bf45-e2793a0d3ea0', 'd1ef8e44-dd8d-4002-9c5a-0520b3fefcfd', 'OPEN', NOW(), NOW(), NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM events WHERE id IN ('d3407173-3984-460e-8429-327b878667ff', 'd1ef8e44-dd8d-4002-9c5a-0520b3fefcfd', '8ce58cdc-3f20-419f-8b61-c3ecf4aa7975');
DELETE FROM seats WHERE id IN ('0a526788-724b-4bf6-9521-56b2afa2a584', '74fffff4-4f46-46e0-8099-14486373296a', '05ded810-dc4b-461b-bf45-e2793a0d3ea0');
-- +goose StatementEnd
