-- +goose Up
-- +goose StatementBegin
create table orders_item (
    id int not null auto_increment,
    name varchar(100) not null,
    price int not null,
    expired_at varchar(20) not null,
    created_at timestamp(3) default current_timestamp(3),
    updated_at timestamp(3) default null,
    deleted_at timestamp(3) default null,
    primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orders_item;
-- +goose StatementEnd
