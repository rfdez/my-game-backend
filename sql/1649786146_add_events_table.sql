CREATE TABLE IF NOT EXISTS events (
  id uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  date DATE NOT NULL,
  keywords VARCHAR(255) ARRAY NOT NULL
);
