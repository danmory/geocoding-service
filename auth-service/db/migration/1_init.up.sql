CREATE TABLE users (
	username varchar PRIMARY KEY,
	password varchar NOT NULL
);

CREATE INDEX ON users(username);
