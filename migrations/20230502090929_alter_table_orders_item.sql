-- +goose Up
-- +goose StatementBegin
alter table orders_item
add column status_deleted bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table orders_item
drop column status_deleted;
-- +goose StatementEnd
