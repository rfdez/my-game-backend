CREATE TABLE IF NOT EXISTS events (
  id uuid PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  date DATE NOT NULL,
  shown INTEGER DEFAULT 0,
  keywords VARCHAR(255) ARRAY NOT NULL,
);

INSERT INTO events (id, name, date, keywords) VALUES (
  'f8b8f8b8-f8b8-f8b8-f8b8-f8b8f8b8f8b8',
  'Event 1',
  '2020-01-01',
  ARRAY['keyword1', 'keyword2']
);
