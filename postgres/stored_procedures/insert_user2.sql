CREATE OR REPLACE PROCEDURE insert_user(_id int, _username text)
LANGUAGE SQL
AS $$
INSERT INTO public.users(
	id, username)
	VALUES ($1, $2);

$$;

CALL insert_user(2, 'jim');