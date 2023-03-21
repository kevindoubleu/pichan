\c habits;

TRUNCATE scorecard;

INSERT INTO
    scorecard (name, connotation, sortOrder)
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
