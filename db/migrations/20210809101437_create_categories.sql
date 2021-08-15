
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    ancestor_id INTEGER REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories;