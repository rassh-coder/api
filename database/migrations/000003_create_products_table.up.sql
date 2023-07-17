CREATE TABLE IF NOT EXISTS products (
    id serial primary key,
    sku text,
    name varchar(255) not null,
    body varchar(255),
    description text,
    price float default 0,
    color varchar(255),
    size ENUM('Large', 'Medium', 'Small'),
    count int default 0
);
