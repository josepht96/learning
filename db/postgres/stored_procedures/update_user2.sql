CREATE OR REPLACE PROCEDURE update_user(_email varchar)
LANGUAGE SQL
AS $$

UPDATE users
SET email = _email
WHERE id = 3;

$$;

CALL update_user('test3@gmail.com');