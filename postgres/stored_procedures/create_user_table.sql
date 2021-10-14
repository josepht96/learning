CREATE OR REPLACE PROCEDURE create_users_table()
LANGUAGE SQL
AS $$
  CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  username text NOT NULL
);

$$;

CALL create_users_table();