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

CREATE TABLE  streets(
id integer PRIMARY KEY ,
name VARCHAR(128) NOT NULL);

CREATE INDEX ON streets (name  varchar_pattern_ops);

-- copy streets(id, name) FROM '/home/ns/tp/bd/back/csv/street.csv' (DELIMITER ';');
-- copy streets(id, name) FROM 'csv/street.csv' (DELIMITER ';');

CREATE TABLE houses (
house VARCHAR(20) NOT NULL,
street_id integer  REFERENCES streets ON DELETE CASCADE) ;
 
CREATE INDEX ON houses (house varchar_pattern_ops);

-- copy houses(street_id, house) FROM '/home/ns/tp/bd/back/csv/houses.csv' (DELIMITER ';');
-- copy houses(street_id, house) FROM 'csv/houses.csv' (DELIMITER ';');

CREATE TYPE order_dish AS (id integer,count int, price int);

CREATE OR REPLACE FUNCTION public.total_cost() RETURNS trigger
AS $$
DECLARE
    rest_id1 integer;
    rest_id_next integer;
    price_per_one integer;
    total_price integer;
    arr order_dish[];
BEGIN
    IF array_length( NEW.cart,1) = 0 THEN
        RETURN NULL;
    END IF;
    SELECT restaurant INTO rest_id1 FROM dish WHERE id=NEW.cart[1].id;
    total_price=0;
    FOR i IN 1..array_length(NEW.cart,1)
    LOOP
        SELECT restaurant,price INTO rest_id_next, price_per_one FROM dish WHERE id=NEW.cart[i].id;
        IF rest_id1 != rest_id_next THEN
            RETURN NULL; 
        END IF;
        total_price = total_price + NEW.cart[i].count*price_per_one;
        arr = array_append(arr, (NEW.cart[i].id, NEW.cart[i].count, price_per_one)::order_dish);
    END LOOP;
    NEW.cart = arr;
    NEW.restaurant_id = rest_id1;
    NEW.total_price = total_price; 
    RAISE NOTICE '%', NEW;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TABLE orders (
id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
restaurant_id integer  REFERENCES restaurants ON DELETE NO ACTION NOT NULL,
date TIMESTAMP DEFAULT now() NOT NULL,
user_id integer  REFERENCES users ON DELETE NO ACTION NOT NULL,
address VARCHAR(256) NOT NULL,
comment VARCHAR(256),
cart order_dish[] NOT NULL,
total_price integer NOT NULL);

CREATE TRIGGER cost_and_restaurant_id BEFORE UPDATE OR INSERT ON orders
 FOR EACH ROW EXECUTE FUNCTION total_cost();