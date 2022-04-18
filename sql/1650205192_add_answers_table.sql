CREATE TABLE IF NOT EXISTS answers (
  event_id uuid REFERENCES events(id),
  question_id uuid REFERENCES questions(id),
  text VARCHAR(255) NOT NULL
);
