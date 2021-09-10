
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
INSERT INTO categories (id, ancestor_id, name) VALUES (9, NULL, 'IT・ガジェット');
INSERT INTO categories (id, ancestor_id, name) VALUES (10, NULL, 'ファッション');
INSERT INTO categories (id, ancestor_id, name) VALUES (11, NULL, '美容');
INSERT INTO categories (id, ancestor_id, name) VALUES (12, NULL, '音楽');
INSERT INTO categories (id, ancestor_id, name) VALUES (13, NULL, '旅行');
INSERT INTO categories (id, ancestor_id, name) VALUES (14, NULL, 'ヘルスケア・フィットネス');
INSERT INTO categories (id, ancestor_id, name) VALUES (15, NULL, 'ビジネス');
INSERT INTO categories (id, ancestor_id, name) VALUES (16, 15, 'ビジネス総合');
INSERT INTO categories (id, ancestor_id, name) VALUES (17, 15, '金融');
INSERT INTO categories (id, ancestor_id, name) VALUES (18, 15, 'IT');
INSERT INTO categories (id, ancestor_id, name) VALUES (19, 15, '不動産');
INSERT INTO categories (id, ancestor_id, name) VALUES (20, 15, '医療・福祉');
INSERT INTO categories (id, ancestor_id, name) VALUES (21, 15, '小売・流通');
INSERT INTO categories (id, ancestor_id, name) VALUES (22, 15, '広告');
INSERT INTO categories (id, ancestor_id, name) VALUES (23, NULL, '車');
INSERT INTO categories (id, ancestor_id, name) VALUES (24, NULL, 'シニア');
INSERT INTO categories (id, ancestor_id, name) VALUES (25, NULL, '就職・転職');
INSERT INTO categories (id, ancestor_id, name) VALUES (26, NULL, 'スポーツ');
INSERT INTO categories (id, ancestor_id, name) VALUES (27, NULL, '暮らし');
INSERT INTO categories (id, ancestor_id, name) VALUES (28, 27, '料理');
INSERT INTO categories (id, ancestor_id, name) VALUES (29, 27, '家事');
INSERT INTO categories (id, ancestor_id, name) VALUES (30, 27, '子育て');
INSERT INTO categories (id, ancestor_id, name) VALUES (31, 27, 'インテリア');
INSERT INTO categories (id, ancestor_id, name) VALUES (32, 27, 'ガーデニング');
INSERT INTO categories (id, ancestor_id, name) VALUES (33, NULL, 'ペット');
INSERT INTO categories (id, ancestor_id, name) VALUES (34, NULL, 'アウトドア');
INSERT INTO categories (id, ancestor_id, name) VALUES (35, 34, 'キャンプ');
INSERT INTO categories (id, ancestor_id, name) VALUES (36, 34, '自転車');
INSERT INTO categories (id, ancestor_id, name) VALUES (37, NULL, '恋愛');
INSERT INTO categories (id, ancestor_id, name) VALUES (38, NULL, '海外');
INSERT INTO categories (id, ancestor_id, name) VALUES (39, NULL, '勉強');
INSERT INTO categories (id, ancestor_id, name) VALUES (40, NULL, 'エンタメ');
INSERT INTO categories (id, ancestor_id, name) VALUES (41, 40, 'アニメ・映画');
INSERT INTO categories (id, ancestor_id, name) VALUES (42, 40, 'ゲーム');
INSERT INTO categories (id, ancestor_id, name) VALUES (43, 40, 'マンガ');
INSERT INTO categories (id, ancestor_id, name) VALUES (44, NULL, 'アート');
INSERT INTO categories (id, ancestor_id, name) VALUES (45, NULL, 'メンズ');
INSERT INTO categories (id, ancestor_id, name) VALUES (46, NULL, 'レディース');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from categories WHERE id BETWEEN 1 AND 46;
SELECT setval('categories_id_seq', 1, false);
