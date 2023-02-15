\c books;

DROP TYPE IF EXISTS BOOK_ACTION CASCADE;

CREATE TYPE BOOK_ACTION AS ENUM ('Bought', 'Start', 'Complete', 'Hold', 'Resume');
