CREATE TABLE IF NOT EXISTS users (
    id              serial             primary key,
    first_name      varchar(50)     not null,
    last_name       varchar(50)     not null,
    username        varchar(50)     not null unique,
    password_hash   varchar(255)    not null
);