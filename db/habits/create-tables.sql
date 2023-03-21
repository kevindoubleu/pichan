\c habits;

DROP TABLE IF EXISTS scorecard;

CREATE TABLE scorecard (
    id SERIAL,
    name VARCHAR(50) NOT NULL,
    connotation HABIT_CONNOTATION NOT NULL DEFAULT 'Neutral',
    time TIME,
    sortOrder INT NOT NULL DEFAULT 1000
);
