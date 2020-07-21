CREATE TABLE fruits
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
INSERT INTO fruits (name)
VALUES ('Orange'),
       ('Pear'),
       ('Apple');