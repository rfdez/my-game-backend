COPY events(id, name, date, keywords)
FROM
  '/docker-entrypoint-initdb.d/events.csv' WITH DELIMITER ',' CSV HEADER;
