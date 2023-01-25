-- psql -U postgres -d postgres -f ./simple_table.
CREATE OR REPLACE PROCEDURE create_projects_table()
LANGUAGE SQL
AS $$
  CREATE TABLE IF NOT EXISTS projects (
  id serial PRIMARY KEY,
  name VARCHAR(50),
  semvar VARCHAR(25),
  description VARCHAR(250),
  timestamp timestamp default current_timestamp
);

$$;

CALL create_projects_table();