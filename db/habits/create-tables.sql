\c habits;

DROP TABLE IF EXISTS scorecards;

CREATE TABLE scorecards (
    id SERIAL,
    name VARCHAR(50) NOT NULL,
    connotation HABIT_CONNOTATION NOT NULL DEFAULT 'Neutral',
    time TIME,
    sortOrder INT NOT NULL DEFAULT 1000
);
