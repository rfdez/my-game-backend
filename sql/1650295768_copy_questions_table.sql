COPY questions(id, text)
FROM
  '/docker-entrypoint-initdb.d/questions.csv' WITH DELIMITER ',' CSV HEADER;
