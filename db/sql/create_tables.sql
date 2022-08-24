\c notes-app

CREATE TABLE IF NOT EXISTS users 
(
    id SERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    passwordhash VARCHAR(100) NOT NULL,
    lastlogin TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deletedat TIMESTAMP
)

CREATE TABLE IF NOT EXISTS notes
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	user_id int NOT NULL,
	created_at timestamp DEFAULT now(),
	updated_at timestamp DEFAULT now(),
	deleted_at timestamp DEFAULT now(),
	CONSTRAINT fk_user
	FOREIGN KEY(user_id)
	REFERENCES users(id)
)
