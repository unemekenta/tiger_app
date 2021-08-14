
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO categories (id, ancestor_id, name) VALUES (1, NULL, 'コンピュータサイエンス');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories WHERE id = 1 OR ancestor_id = 1;
