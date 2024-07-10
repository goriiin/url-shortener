create table if not exists url
(
    id serial primary key,
    alias text NOT NULL unique,
    url text NOT NULL UNIQUE,
    created_at timestamp not null
);