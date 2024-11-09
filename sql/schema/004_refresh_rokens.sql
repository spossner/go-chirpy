-- +goose Up
create table REFRESH_TOKENS (
    token varchar primary key,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    user_id uuid not null references users(id) on delete cascade,
    expires_at timestamp not null,
    revoked_at timestamp
);

-- +goose Down
drop table REFRESH_TOKENS;