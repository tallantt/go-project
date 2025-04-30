CREATE TABLE cars (
                      id SERIAL PRIMARY KEY,
                      brand TEXT NOT NULL,
                      model TEXT NOT NULL UNIQUE,
                      year TEXT NOT NULL,
                      color TEXT NOT NULL
);