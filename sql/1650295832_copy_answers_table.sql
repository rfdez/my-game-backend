COPY answers(event_id, question_id, text)
FROM
  '/docker-entrypoint-initdb.d/answers.csv' WITH DELIMITER ',' CSV HEADER;
