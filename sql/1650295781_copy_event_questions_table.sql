COPY event_questions(event_id, question_id, round)
FROM
  '/docker-entrypoint-initdb.d/event_questions.csv' WITH DELIMITER ',' CSV HEADER;
