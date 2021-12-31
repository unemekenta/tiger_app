
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO categories_websites (category_id, website_id) VALUES (17, 1);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 2);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 3);
INSERT INTO categories_websites (category_id, website_id) VALUES (46, 4);
INSERT INTO categories_websites (category_id, website_id) VALUES (46, 5);
INSERT INTO categories_websites (category_id, website_id) VALUES (28, 6);
INSERT INTO categories_websites (category_id, website_id) VALUES (28, 7);
INSERT INTO categories_websites (category_id, website_id) VALUES (55, 8);
INSERT INTO categories_websites (category_id, website_id) VALUES (36, 9);
INSERT INTO categories_websites (category_id, website_id) VALUES (36, 10);
INSERT INTO categories_websites (category_id, website_id) VALUES (38, 11);
INSERT INTO categories_websites (category_id, website_id) VALUES (52, 12);
INSERT INTO categories_websites (category_id, website_id) VALUES (47, 13);
INSERT INTO categories_websites (category_id, website_id) VALUES (47, 14);
INSERT INTO categories_websites (category_id, website_id) VALUES (47, 15);
INSERT INTO categories_websites (category_id, website_id) VALUES (45, 16);
INSERT INTO categories_websites (category_id, website_id) VALUES (48, 17);
INSERT INTO categories_websites (category_id, website_id) VALUES (49, 18);
INSERT INTO categories_websites (category_id, website_id) VALUES (31, 19);
INSERT INTO categories_websites (category_id, website_id) VALUES (31, 20);
INSERT INTO categories_websites (category_id, website_id) VALUES (31, 21);
INSERT INTO categories_websites (category_id, website_id) VALUES (40, 22);
INSERT INTO categories_websites (category_id, website_id) VALUES (40, 23);
INSERT INTO categories_websites (category_id, website_id) VALUES (40, 24);
INSERT INTO categories_websites (category_id, website_id) VALUES (50, 25);
INSERT INTO categories_websites (category_id, website_id) VALUES (50, 26);
INSERT INTO categories_websites (category_id, website_id) VALUES (50, 27);
INSERT INTO categories_websites (category_id, website_id) VALUES (53, 28);
INSERT INTO categories_websites (category_id, website_id) VALUES (26, 29);
INSERT INTO categories_websites (category_id, website_id) VALUES (41, 30);
INSERT INTO categories_websites (category_id, website_id) VALUES (34, 31);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 32);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 33);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 34);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 35);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 36);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 37);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 38);
INSERT INTO categories_websites (category_id, website_id) VALUES (17, 39);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 40);
INSERT INTO categories_websites (category_id, website_id) VALUES (28, 41);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 42);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 43);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 44);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 45);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 46);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 47);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 48);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 49);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 50);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 51);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 52);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 53);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 54);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 55);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 56);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 57);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 58);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 59);
INSERT INTO categories_websites (category_id, website_id) VALUES (23, 60);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 61);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 62);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 63);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 64);
INSERT INTO categories_websites (category_id, website_id) VALUES (17, 65);
INSERT INTO categories_websites (category_id, website_id) VALUES (48, 66);
INSERT INTO categories_websites (category_id, website_id) VALUES (35, 67);
INSERT INTO categories_websites (category_id, website_id) VALUES (35, 68);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 69);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 70);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 71);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 72);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 73);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 74);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 75);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 76);
INSERT INTO categories_websites (category_id, website_id) VALUES (24, 77);
INSERT INTO categories_websites (category_id, website_id) VALUES (22, 78);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 79);
INSERT INTO categories_websites (category_id, website_id) VALUES (22, 80);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 81);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 82);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 83);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 84);
INSERT INTO categories_websites (category_id, website_id) VALUES (16, 85);
INSERT INTO categories_websites (category_id, website_id) VALUES (18, 86);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories_websites WHERE id BETWEEN 1 AND 86;
SELECT setval('categories_websites_id_seq', 1, false);
;