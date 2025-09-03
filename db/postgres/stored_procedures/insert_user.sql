CREATE OR REPLACE PROCEDURE insert_user(_id int, _username text)
LANGUAGE SQL
AS $$
INSERT INTO public.users(
	id, username)
	VALUES (_id, _username);

$$;

CALL insert_user(3, 'jake');