CREATE TABLE collections (
    id   serial primary key,
    name varchar(255) not null unique
);

CREATE TABLE users (
    id            serial primary key,
    login         varchar(255) not null unique,
    password_hash varchar(255) not null,
    active        boolean default true not null,
    collection_id int references collections(id) on delete cascade
);

CREATE TABLE api_keys (
    id        serial primary key,
    user_id   int not null references users(id) on delete cascade,
    key       text not null unique
);

CREATE UNIQUE INDEX idx_api_keys_key ON api_keys(key);