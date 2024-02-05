CREATE TABLE books (
    ID SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, description TEXT, author VARCHAR(255), published_at TIMESTAMP
);