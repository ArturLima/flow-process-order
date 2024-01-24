create table customer(
 id UUID primary key,
 name VARCHAR(20),
 age INT,
 created_at DATE
);

create table product(
id uuid primary key,
name varchar(30),
price decimal,
validate_date DATE,
amount int,
created_ad date
)


create table current_account(
id uuid primary key,
customer_id uuid REFERENCES customer(id),
balance decimal,
created_at date
)

create table "order"(
id uuid primary key,
product_id uuid references product(id),
customer_id uuid references customer(id),
amount int,
created_at date
)

create table order_status(
 order_id uuid references "order"(id),
 status int references STATUS(status_id)
)

create table STATUS(
 status_id int primary key,
 description varchar(10)
)


create table order_log(
order_id uuid references "order"(id),
description varchar(50),
status varchar(50)
)
