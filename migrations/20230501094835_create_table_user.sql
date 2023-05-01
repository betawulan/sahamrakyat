-- +goose Up
-- +goose StatementBegin
create table user (
    id int not null auto_increment,
    fullname varchar(100) not null,
    first_order varchar(100) not null,
    created_at timestamp(3) default current_timestamp(3),
    updated_at timestamp(3) default null,
    deleted_at timestamp(3) default null,
    primary key(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
