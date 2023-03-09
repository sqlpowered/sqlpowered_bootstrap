create table if not exists product (
    id serial primary key,
    name varchar,
    price float
);

create table if not exists city (
    id serial primary key,
    name varchar
);

create table if not exists district (
    id serial primary key,
    name varchar,
    city int references city(id) ON DELETE CASCADE
);

create table if not exists retail_shop (
    id serial primary key,
    name varchar,
    district_id int references district(id) ON DELETE CASCADE
);

create table if not exists retail_shop_product_stock (
    id serial primary key,
    product_id int references product(id) ON DELETE CASCADE,
    retail_shop_id int references retail_shop(id) ON DELETE CASCADE,
    number_stocked int
);


insert into product(price, name)
values
    (12.02, 'hair dryer'),
    (5.22, 'shoe polish'),
    (500.00, 'games console'),
    (1500.00, 'gaming pc');

insert into city(name) 
values
    ('capital'),
    ('remote');


insert into district(name, city) 
values
    ('north', 1),
    ('south', 1),
    ('east', 1),
    ('west', 1),
    ('north', 2),
    ('south', 2),
    ('east', 2),
    ('west', 2);

insert into retail_shop(name, district_id) 
values
    ('central', 1),
    ('nice end of town', 1),
    ('less wonderful part of town', 1),
    ('middle of nowhere', 5);

insert into retail_shop_product_stock (product_id, retail_shop_id, number_stocked)
values
    (1, 1, 5),
    (1, 2, 5),
    (1, 3, 5),
    (1, 4, 50);