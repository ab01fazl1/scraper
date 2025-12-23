-- +goose UP
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    creted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose DOWN
DROP TABLE users;