\c habits;

TRUNCATE scorecards;

INSERT INTO
    scorecards (name, connotation, sortOrder)
VALUES
    (
        'wake up',
        'Neutral',
        500
    ),
    (
        'scroll ig',
        'Negative',
        1000
    ),
    (
        'making the bed',
        'Positive',
        1001
    );
