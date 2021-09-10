
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO categories_websites (category_id, website_id) VALUES (17, 1);
INSERT INTO categories_websites (category_id, website_id) VALUES (9, 2);
INSERT INTO categories_websites (category_id, website_id) VALUES (9, 3);
INSERT INTO categories_websites (category_id, website_id) VALUES (10, 4);
INSERT INTO categories_websites (category_id, website_id) VALUES (10, 5);
INSERT INTO categories_websites (category_id, website_id) VALUES (27, 6);
INSERT INTO categories_websites (category_id, website_id) VALUES (27, 7);
INSERT INTO categories_websites (category_id, website_id) VALUES (44, 8);
INSERT INTO categories_websites (category_id, website_id) VALUES (26, 9);
INSERT INTO categories_websites (category_id, website_id) VALUES (26, 10);
INSERT INTO categories_websites (category_id, website_id) VALUES (26, 11);
INSERT INTO categories_websites (category_id, website_id) VALUES (12, 12);
INSERT INTO categories_websites (category_id, website_id) VALUES (45, 13);
INSERT INTO categories_websites (category_id, website_id) VALUES (45, 14);
INSERT INTO categories_websites (category_id, website_id) VALUES (45, 15);
INSERT INTO categories_websites (category_id, website_id) VALUES (40, 16);
INSERT INTO categories_websites (category_id, website_id) VALUES (46, 17);
INSERT INTO categories_websites (category_id, website_id) VALUES (11, 18);
INSERT INTO categories_websites (category_id, website_id) VALUES (30, 19);
INSERT INTO categories_websites (category_id, website_id) VALUES (30, 20);
INSERT INTO categories_websites (category_id, website_id) VALUES (30, 21);
INSERT INTO categories_websites (category_id, website_id) VALUES (13, 22);
INSERT INTO categories_websites (category_id, website_id) VALUES (13, 23);
INSERT INTO categories_websites (category_id, website_id) VALUES (13, 24);
INSERT INTO categories_websites (category_id, website_id) VALUES (14, 25);
INSERT INTO categories_websites (category_id, website_id) VALUES (14, 26);
INSERT INTO categories_websites (category_id, website_id) VALUES (14, 27);
INSERT INTO categories_websites (category_id, website_id) VALUES (25, 28);
INSERT INTO categories_websites (category_id, website_id) VALUES (23, 29);
INSERT INTO categories_websites (category_id, website_id) VALUES (35, 30);
INSERT INTO categories_websites (category_id, website_id) VALUES (33, 31);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories WHERE id BETWEEN 1 AND 33;
SELECT setval('categories_id_seq', 1, false);
