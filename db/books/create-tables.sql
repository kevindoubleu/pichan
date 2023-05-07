\c books;

DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id SERIAL,
    name VARCHAR(50) NOT NULL,
    action BOOK_ACTION NOT NULL,
    time TIMESTAMP NOT NULL DEFAULT NOW()
);
