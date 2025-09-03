CREATE OR REPLACE PROCEDURE update_user()
LANGUAGE SQL
AS $$

UPDATE users
SET email = 'test2@gmail.com'
WHERE id = 3;

$$;

CALL update_user();