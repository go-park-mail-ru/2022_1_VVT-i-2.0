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


-- CREATE TABLE  msc_streets(
-- id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
-- name VARCHAR(256) NOT NULL,
-- code NUMERIC(15) NOT NULL UNIQUE);

-- CREATE TABLE msc_houses (
-- houses VARCHAR(25)[] NOT NULL,
-- street_code NUMERIC(15) NOT NULL,
-- street_id integer  REFERENCES msc_streets ON DELETE CASCADE) ;

-- /my_windows_d/помойка/csv/doma.csv
-- copy msc_streets(name,code) FROM '/home/ns/tp/bd/sugg/street_uniq.csv' (DELIMITER ',');
-- copy msc_houses(houses,street_code) FROM '/home/ns/tp/bd/sugg/doma.csv' (DELIMITER ';');