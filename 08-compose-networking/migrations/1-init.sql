CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title VARCHAR NOT NULL,
	completed BOOLEAN NOT NULL
);

INSERT INTO todos(title, completed) 
	VALUES	
		('Finish the go API for the docker compose test', True),
		('Finish the docker course', false);