-- psql -U postgres -d postgres -f ./simple_table.
CREATE OR REPLACE PROCEDURE create_users_table()
LANGUAGE SQL
AS $$
  CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  name VARCHAR(50),
  timestamp timestamp default current_timestamp
);

$$;

CALL create_users_table();