CREATE TABLE IF NOT EXISTS answers (
  id uuid PRIMARY KEY,
  text VARCHAR(255) NOT NULL,
  question_id uuid REFERENCES questions(id)
);

INSERT INTO
  answers (id, text, question_id)
VALUES
  (
    '2a82fe2b-4a38-4ae2-998d-aa0b961cdf63',
    'Sí',
    (
      SELECT
        id
      FROM
        questions
      WHERE
        text = '¿Sucedió en España?'
    )
  );

INSERT INTO
  answers (id, text, question_id)
VALUES
  (
    '54050f31-6bc1-4bac-8ab6-181897f86627',
    'NO',
    (
      SELECT
        id
      FROM
        questions
      WHERE
        text = '¿Sucedió en un ámbito internacional?'
    )
  );

INSERT INTO
  answers (id, text, question_id)
VALUES
  (
    '87175ffe-c762-4103-8fc9-cec652ac2d82',
    'NO',
    (
      SELECT
        id
      FROM
        questions
      WHERE
        text = '¿Sucedió en América?'
    )
  );

INSERT INTO
  answers (id, text, question_id)
VALUES
  (
    '479885d5-87ff-4a80-a4dd-5a8daa176cb1',
    'NO',
    (
      SELECT
        id
      FROM
        questions
      WHERE
        text = '¿Fué internacionalmente relevante?'
    )
  );
