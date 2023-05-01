-- +goose Up
-- +goose StatementBegin
alter table user
add column status_deleted bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table user
drop column status_deleted;
-- +goose StatementEnd
