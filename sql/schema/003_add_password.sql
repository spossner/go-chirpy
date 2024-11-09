-- +goose Up
alter table USERS
add hashed_password varchar not null default md5(random()::text);

-- +goose Down
alter table USERS
drop column hashed_password;