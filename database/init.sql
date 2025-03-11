CREATE TABLE roles
(
    role_id   SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE users
(
    user_id       SERIAL PRIMARY KEY,
    username      VARCHAR(255) UNIQUE NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    phone_number  VARCHAR(32)         NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    role_id       INT                 NOT NULL REFERENCES roles (role_id)
);
CREATE TABLE stores
(
    store_id      SERIAL PRIMARY KEY,
    store_name    VARCHAR(255) NOT NULL UNIQUE,
    store_address TEXT,
    phone_number  VARCHAR(20)
);
CREATE TABLE categories
(
    category_id          SERIAL PRIMARY KEY,
    category_description TEXT,
    category_name        VARCHAR(50) NOT NULL UNIQUE
);
CREATE TABLE manufacturers
(
    manufacturer_id   SERIAL PRIMARY KEY,
    manufacturer_name VARCHAR(255) NOT NULL
);

CREATE TABLE instruments
(
    instrument_id   SERIAL PRIMARY KEY,
    instrument_name VARCHAR(255)   NOT NULL,
    category_id     INT            NOT NULL REFERENCES categories (category_id),
    manufacturer_id INT            NOT NULL REFERENCES manufacturers (manufacturer_id),
    store_id        INT            NOT NULL REFERENCES stores (store_id),
    description     TEXT,
    price_per_day   DECIMAL(10, 2) NOT NULL,
    image_url       TEXT           -- Добавляем новый столбец, NULL по умолчанию
);
CREATE TABLE rentals
(
    rental_id     SERIAL PRIMARY KEY,
    user_id       INT       NOT NULL REFERENCES users (user_id),
    instrument_id INT       NOT NULL REFERENCES instruments (instrument_id),
    rental_date   TIMESTAMP NOT NULL,
    return_date   TIMESTAMP
);
CREATE TABLE payments
(
    payment_id     SERIAL PRIMARY KEY,
    rental_id      INT            NOT NULL REFERENCES rentals (rental_id),
    payment_date   TIMESTAMP      NOT NULL,
    payment_amount DECIMAL(10, 2) NOT NULL CHECK (payment_amount >= 0)
);
CREATE TABLE repairs
(
    repair_id         SERIAL PRIMARY KEY,
    instrument_id     INT       NOT NULL REFERENCES instruments (instrument_id),
    repair_start_date TIMESTAMP NOT NULL,
    repair_end_date   TIMESTAMP NOT NULL,
    repair_cost       DECIMAL(10, 2) CHECK (repair_cost >= 0),
    description       TEXT
);
CREATE TABLE discounts
(
    discount_id         SERIAL PRIMARY KEY,
    instrument_id       INT       NOT NULL REFERENCES instruments (instrument_id),
    discount_percentage DECIMAL(5, 2) CHECK (discount_percentage BETWEEN 0 AND 100),
    valid_until         TIMESTAMP NOT NULL
);
CREATE TABLE reviews
(
    review_id   SERIAL PRIMARY KEY,
    rental_id   INT NOT NULL REFERENCES rentals (rental_id),
    review_text TEXT,
    rating      INT CHECK (rating BETWEEN 1 AND 5)
);

-- Добавляем роли пользователей
INSERT INTO roles (role_name)
VALUES ('customer'),
       ('staff'),
       ('chief'),
       ('admin');

-- Добавляем магазины
INSERT INTO stores (store_name, store_address, phone_number)
VALUES ('Downtown Store', '123 Main St', '+123456789'),
       ('Uptown Store', '456 Elm St', '+987654321');

-- Добавляем категории инструментов
INSERT INTO categories (category_name, category_description) VALUES 
    ('Струнные', 'Струнные музыкальные инструменты'),
    ('Духовые', 'Духовые музыкальные инструменты'),
    ('Клавишные', 'Клавишные музыкальные инструменты'),
    ('Ударные', 'Ударные музыкальные инструменты'),
    ('Электронные', 'Электронные музыкальные инструменты');

-- Добавляем производителей
INSERT INTO manufacturers (manufacturer_name) VALUES 
    ('Yamaha'),
    ('Roland'),
    ('Fender'),
    ('Gibson'),
    ('Pearl'),
    ('Ibanez'),
    ('Casio'),
    ('Kawai');
