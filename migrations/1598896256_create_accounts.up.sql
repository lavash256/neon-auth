CREATE TABLE accounts (
    id bigserial not null primary key,
    email varchar not null unique,
    password varchar not null,
    status varchar not null,
    created_at time not null
)