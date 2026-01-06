-- +goose Up
CREATE TABLE users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	preferences TEXT,
	mail TEXT,
	phone_number TEXT,
	CHECK (mail IS NOT NULL OR phone_number IS NOT NULL)
);
-- +goose Down
DROP TABLE users;
