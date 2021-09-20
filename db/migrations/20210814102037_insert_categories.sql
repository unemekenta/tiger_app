
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO categories (id, ancestor_id, name) VALUES (1, NULL, 'プロダクト');
INSERT INTO categories (id, ancestor_id, name) VALUES (2, NULL, 'ファッション');
INSERT INTO categories (id, ancestor_id, name) VALUES (3, NULL, '美容・健康');
INSERT INTO categories (id, ancestor_id, name) VALUES (4, NULL, '音楽');
INSERT INTO categories (id, ancestor_id, name) VALUES (5, NULL, 'エンタメ');
INSERT INTO categories (id, ancestor_id, name) VALUES (6, NULL, 'ビジネス');
INSERT INTO categories (id, ancestor_id, name) VALUES (7, NULL, 'スポーツ');
INSERT INTO categories (id, ancestor_id, name) VALUES (8, NULL, '投資');
INSERT INTO categories (id, ancestor_id, name) VALUES (9, NULL, '暮らし');
INSERT INTO categories (id, ancestor_id, name) VALUES (10, NULL, 'アウトドア');
INSERT INTO categories (id, ancestor_id, name) VALUES (11, NULL, 'グルメ');
INSERT INTO categories (id, ancestor_id, name) VALUES (12, NULL, '恋愛・結婚');
INSERT INTO categories (id, ancestor_id, name) VALUES (13, NULL, '学び');
INSERT INTO categories (id, ancestor_id, name) VALUES (14, NULL, '就職・転職');
INSERT INTO categories (id, ancestor_id, name) VALUES (15, NULL, 'デザイン・アート');
INSERT INTO categories (id, ancestor_id, name) VALUES (16, 6, 'ビジネス総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (17, 6, '金融');
INSERT INTO categories (id, ancestor_id, name) VALUES (18, 6, 'IT');
INSERT INTO categories (id, ancestor_id, name) VALUES (19, 6, '不動産');
INSERT INTO categories (id, ancestor_id, name) VALUES (20, 6, '医療・福祉');
INSERT INTO categories (id, ancestor_id, name) VALUES (21, 6, '小売・流通');
INSERT INTO categories (id, ancestor_id, name) VALUES (22, 6, '広告');
INSERT INTO categories (id, ancestor_id, name) VALUES (23, 6, '人材・HR');
INSERT INTO categories (id, ancestor_id, name) VALUES (24, 6, 'マーケティング');
INSERT INTO categories (id, ancestor_id, name) VALUES (25, 1, 'ガジェット');
INSERT INTO categories (id, ancestor_id, name) VALUES (26, 1, '車');
INSERT INTO categories (id, ancestor_id, name) VALUES (27, 1, '自転車');
INSERT INTO categories (id, ancestor_id, name) VALUES (28, 9, '暮らし総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (29, 9, '料理');
INSERT INTO categories (id, ancestor_id, name) VALUES (30, 9, '家事');
INSERT INTO categories (id, ancestor_id, name) VALUES (31, 9, '子育て');
INSERT INTO categories (id, ancestor_id, name) VALUES (32, 9, 'インテリア');
INSERT INTO categories (id, ancestor_id, name) VALUES (33, 9, 'ガーデニング');
INSERT INTO categories (id, ancestor_id, name) VALUES (34, 9, 'ペット');
INSERT INTO categories (id, ancestor_id, name) VALUES (35, 9, 'お金');
INSERT INTO categories (id, ancestor_id, name) VALUES (36, 7, 'スポーツ総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (37, 7, '野球');
INSERT INTO categories (id, ancestor_id, name) VALUES (38, 7, 'バスケ');
INSERT INTO categories (id, ancestor_id, name) VALUES (39, 7, 'サッカー');
INSERT INTO categories (id, ancestor_id, name) VALUES (40, 10, '旅行');
INSERT INTO categories (id, ancestor_id, name) VALUES (41, 10, 'キャンプ');
INSERT INTO categories (id, ancestor_id, name) VALUES (42, 5, 'アニメ・映画');
INSERT INTO categories (id, ancestor_id, name) VALUES (43, 5, 'ゲーム');
INSERT INTO categories (id, ancestor_id, name) VALUES (44, 5, 'マンガ');
INSERT INTO categories (id, ancestor_id, name) VALUES (45, 5, '芸能');
INSERT INTO categories (id, ancestor_id, name) VALUES (46, 2, 'ファッション総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (47, 2, 'メンズ');
INSERT INTO categories (id, ancestor_id, name) VALUES (48, 2, 'レディース');
INSERT INTO categories (id, ancestor_id, name) VALUES (49, 3, '美容');
INSERT INTO categories (id, ancestor_id, name) VALUES (50, 3, 'ヘルスケア・フィットネス');
INSERT INTO categories (id, ancestor_id, name) VALUES (51, 13, '勉強');
INSERT INTO categories (id, ancestor_id, name) VALUES (52, 4, '音楽総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (53, 14, '就職・転職総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (54, 8, '投資総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (55, 15, 'プロダクトデザイン');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories WHERE id BETWEEN 1 AND 55;
SELECT setval('categories_id_seq', 1, false);
