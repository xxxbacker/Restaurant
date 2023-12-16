CREATE TABLE IF NOT EXISTS account (
    account_id SERIAL PRIMARY KEY,
    post varchar(50) not null,
    nickname VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL
    );

CREATE TABLE IF NOT EXISTS courier (
    courier_id serial primary key ,
    title varchar(50) not null ,
    phone varchar(50) not null
    );

CREATE TABLE IF NOT EXISTS ord (
    ord_id SERIAL PRIMARY KEY,
    ord_date date NOT NULL,
    created_at timestamp NOT null,
    account_id BIGINT references account(account_id),
    courier_id BIGINT references courier(courier_id)
    );

CREATE TABLE IF NOT EXISTS cheque (
    cheque_id serial primary key ,
    price int not null,
    pay_method varchar(50) not null,
    pay_status varchar(50) not null,
    created_at timestamp not null,
    ord_id BIGINT references ord(ord_id)
    );

CREATE TABLE IF NOT EXISTS menu_item (
    menu_id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    price int not null ,
    created_at TIMESTAMP NOT null
);

CREATE TABLE IF NOT EXISTS order_item (
    order_item_id SERIAL PRIMARY KEY,
    title varchar(50),
    price int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    menu_id BIGINT references menu_item(menu_id),
    ord_id BIGINT references ord(ord_id)
);
