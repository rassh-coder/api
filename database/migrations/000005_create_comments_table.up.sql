CREATE TABLE IF NOT EXISTS comments (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    body text not null,
    rate float not null,
    product_id bigint unsigned not null,
    foreign key (product_id) references products(id) on delete cascade
);