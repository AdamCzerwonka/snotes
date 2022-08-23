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
