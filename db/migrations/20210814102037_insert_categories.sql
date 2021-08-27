
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO categories (id, ancestor_id, name) VALUES (1, NULL, 'コンピュータサイエンス');
INSERT INTO categories (id, ancestor_id, name) VALUES (2, NULL, 'web開発');
INSERT INTO categories (id, ancestor_id, name) VALUES (3, 1, 'アルゴリズム');
INSERT INTO categories (id, ancestor_id, name) VALUES (4, 1, 'ネットワーク');
INSERT INTO categories (id, ancestor_id, name) VALUES (5, 1, 'セキュリティ');
INSERT INTO categories (id, ancestor_id, name) VALUES (6, 2, 'フロントエンド');
INSERT INTO categories (id, ancestor_id, name) VALUES (7, 2, 'バックエンド');
INSERT INTO categories (id, ancestor_id, name) VALUES (8, 2, 'インフラ');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories WHERE id = 1 OR ancestor_id = 1;
DELETE from categories WHERE id = 2 OR ancestor_id = 2;
