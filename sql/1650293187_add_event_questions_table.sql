CREATE TABLE IF NOT EXISTS event_questions (
  event_id uuid REFERENCES events(id),
  question_id uuid REFERENCES questions(id),
  round INTEGER NOT NULL
);
