CREATE TABLE IF NOT EXISTS events (
  id uuid PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  date DATE NOT NULL,
  shown INTEGER DEFAULT 0,
  keywords VARCHAR(255) ARRAY NOT NULL
);

INSERT INTO
  events (id, name, date, keywords)
VALUES
  (
    'f8b8f8b8-f8b8-f8b8-f8b8-f8b8f8b8f8b8',
    'La República',
    '2022-04-14',
    ARRAY ['republica', 'españa', 'española']
  );
