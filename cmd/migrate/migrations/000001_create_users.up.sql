CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key,
    email citext unique not null,
    username varchar(255) unique not null,
    password bytea not null,
    created_at timestamptz not null default now(),
);