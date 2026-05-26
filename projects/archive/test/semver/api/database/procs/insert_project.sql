CREATE OR REPLACE PROCEDURE insert_project(_id int, _name text, _semvar VARCHAR, _description VARCHAR)
LANGUAGE SQL
AS $$
INSERT INTO public.projects(
	id, name, semvar, description
    )
	VALUES ($1, $2, $3, $4);

$$;

CALL insert_project(0, 'test_semvar', '0.0.0', 'a test semvar entry');