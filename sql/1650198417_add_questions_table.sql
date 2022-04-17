CREATE TABLE IF NOT EXISTS questions (
  id uuid PRIMARY KEY,
  text VARCHAR(255) NOT NULL,
  round INTEGER NOT NULL,
  event_id uuid REFERENCES events(id)
);

INSERT INTO
  questions (id, text, round, event_id)
VALUES
  (
    'f8b8f8b8-f8b8-f8b8-f8b8-f8b8f8b8f8b8',
    '¿Sucedió en España?',
    1,
    (
      SELECT
        id
      FROM
        events
      WHERE
        name = 'La República'
    )
  );

INSERT INTO
  questions (id, text, round, event_id)
VALUES
  (
    'b4eaeb59-d5df-4255-bb9c-2b74f9df6a63',
    '¿Sucedió en un ámbito internacional?',
    1,
    (
      SELECT
        id
      FROM
        events
      WHERE
        name = 'La República'
    )
  );

INSERT INTO
  questions (id, text, round, event_id)
VALUES
  (
    '2c983429-4d9b-4917-8a80-d519a098d161',
    '¿Sucedió en América?',
    1,
    (
      SELECT
        id
      FROM
        events
      WHERE
        name = 'La República'
    )
  );

INSERT INTO
  questions (id, text, round, event_id)
VALUES
  (
    '54757601-d61b-4caf-931b-5b1225c5c137',
    '¿Fué internacionalmente relevante?',
    1,
    (
      SELECT
        id
      FROM
        events
      WHERE
        name = 'La República'
    )
  );
