DROP TABLE IF EXISTS "Restaurants";
DROP TABLE IF EXISTS "Dish";
DROP TABLE IF EXISTS "User";


create table "Restaurants"
(
    id INTEGER PRIMARY KEY,
    name varchar(80),
    city varchar(80),
    address text,
    street text,
    house_number int,
    corpus varchar(10),
    flat_office int,
    image_path text,
    price int,
    rating float4,
    time_to_delivery int,
    start_time_delivery int,
    finish_time_delivery int
);

create table "Dish"
(
    id integer primary key,
    restaurant integer references "Restaurants",
    name varchar(80),
    description text,
    image text,
    calories integer,
    price integer
);

CREATE TABLE users(
id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
name VARCHAR(256) NOT NULL,
email VARCHAR(256)  CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$') NOT NULL UNIQUE,
phone NUMERIC(11) CHECK(phone>79000000000 and phone <80000000000) NOT NULL UNIQUE  
);

INSERT INTO users(name,email,phone) VALUES
('Наташа','nat-s.skv@mail.ru',79015020456),
('Кирилл','katashinsky-k@yandex.ru',79040666020),
('Андрей','diakonovA@gmail.com',79877434370);
