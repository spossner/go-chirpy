-- +goose Up
create table CHIRPS (
    id          uuid primary key default gen_random_uuid(),
    body        varchar not null,
    user_id     uuid not null references users(id) on delete cascade,
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp
);

-- +goose Down
drop table CHIRPS;