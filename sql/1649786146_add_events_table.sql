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
    'edfc45ac-bc28-46fa-bf32-555725781964',
    'La República',
    (
      SELECT
        CURRENT_DATE
    ),
    ARRAY ['republica', 'españa', 'española']
  );
