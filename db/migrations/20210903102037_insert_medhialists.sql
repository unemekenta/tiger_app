
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO websites (id, name, url, company_name) VALUES (1, 'ZUU online', 'https://zuuonline.com/', '株式会社ZUU');
INSERT INTO websites (id, name, url, company_name) VALUES (2, 'ferret（フェレット）', 'https://ferret-plus.com/', '株式会社ベーシック');
INSERT INTO websites (id, name, url, company_name) VALUES (3, 'バズ部', 'https://bazubu.com/', '株式会社ルーシー');
INSERT INTO websites (id, name, url, company_name) VALUES (4, 'FASHION PRESS（ファッションプレス）', 'https://www.fashion-press.net/', '株式会社カーリン');
INSERT INTO websites (id, name, url, company_name) VALUES (5, 'HYPEBEAST（ハイプビースト）', 'https://hypebeast.com/jp', 'Hypebeast Ltd.');
INSERT INTO websites (id, name, url, company_name) VALUES (6, '北欧、暮らしの道具店', 'https://hokuohkurashi.com/', '株式会社クラシコム');
INSERT INTO websites (id, name, url, company_name) VALUES (7, 'MONOCO（モノコ）', 'https://monoco.jp/', '株式会社MONOCO');
INSERT INTO websites (id, name, url, company_name) VALUES (8, 'AXIS（アクシス）', 'https://www.axismag.jp/', '株式会社アクシス');
INSERT INTO websites (id, name, url, company_name) VALUES (9, 'VICTORY（ビクトリー）', 'https://victorysportsnews.com/', '株式会社VICTORY');
INSERT INTO websites (id, name, url, company_name) VALUES (10, 'REALSPORTS（リアルスポーツ）', 'https://real-sports.jp/', '株式会社 REAL SPORTS');
INSERT INTO websites (id, name, url, company_name) VALUES (11, 'BASKETBALLKING（バスケットボールキング）', 'https://basketballking.jp/', '株式会社フロムワン');
INSERT INTO websites (id, name, url, company_name) VALUES (12, 'cinra（シンラ）', 'https://www.cinra.net/', '株式会社cinra');
INSERT INTO websites (id, name, url, company_name) VALUES (13, 'GQ', 'https://www.gqjapan.jp/', '合同会社コンデナスト・ジャパン');
INSERT INTO websites (id, name, url, company_name) VALUES (14, 'TASCLAP（タスクラップ）', 'https://mens.tasclap.jp/', '株式会社カカクコム');
INSERT INTO websites (id, name, url, company_name) VALUES (15, 'Smartlog（スマートログ）', 'https://smartlog.jp/', '株式会社Smartlog');
INSERT INTO websites (id, name, url, company_name) VALUES (16, 'modelpress（モデルプレス）', 'https://mdpr.jp/', '株式会社ネットネイティブ');
INSERT INTO websites (id, name, url, company_name) VALUES (17, 'マイナビウーマン', 'https://woman.mynavi.jp/', '株式会社マイナビ');
INSERT INTO websites (id, name, url, company_name) VALUES (18, '美的.com', 'https://www.biteki.com/', '株式会社小学館');
INSERT INTO websites (id, name, url, company_name) VALUES (19, 'HugKum（はぐくむ）', 'https://hugkum.sho.jp/', '株式会社小学館');
INSERT INTO websites (id, name, url, company_name) VALUES (20, 'kodomoe（コドモエ）', 'https://kodomoe.net/', '株式会社白泉社');
INSERT INTO websites (id, name, url, company_name) VALUES (21, 'KIDSNA（キズナ）', 'https://kidsna.com/magazine', '株式会社ネクストビート');
INSERT INTO websites (id, name, url, company_name) VALUES (22, 'ことりっぷWEB', 'https://co-trip.jp/', '株式会社昭文社ホールディングス');
INSERT INTO websites (id, name, url, company_name) VALUES (23, 'たびらい沖縄', 'https://www.tabirai.net/s/sightseeing/okinawa/', '株式会社パム');
INSERT INTO websites (id, name, url, company_name) VALUES (24, 'tsunagu Japan', 'https://www.tsunagujapan.com/', '株式会社D2CX');
INSERT INTO websites (id, name, url, company_name) VALUES (25, 'FiNc U（フィンクユー）', 'https://u.finc.com', '株式会社 FiNC Technologies');
INSERT INTO websites (id, name, url, company_name) VALUES (26, 'MELOS（メロス）', 'https://melos.media/', '株式会社APPY');
INSERT INTO websites (id, name, url, company_name) VALUES (27, 'Ccury（キュリー）', 'https://cury.jp/media/', '株式会社UOCC');
INSERT INTO websites (id, name, url, company_name) VALUES (28, 'Career Supli（キャリアサプリ）', 'https://careersupli.jp/', 'アデコ株式会社');
INSERT INTO websites (id, name, url, company_name) VALUES (29, 'CarMe（カーミー）', 'https://car-me.jp/', '株式会社ファブリカコミュニケーションズ');
INSERT INTO websites (id, name, url, company_name) VALUES (30, 'TAKIBI（タキビ）', 'https://www.takibi-reservation.style/media/', '株式会社フォーイット');
INSERT INTO websites (id, name, url, company_name) VALUES (31, 'いぬのきもち', 'https://dog.benesse.ne.jp/', '株式会社ベネッセコーポレーション');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from websites WHERE id BETWEEN 1 AND 38;
SELECT setval('websites_id_seq', 1, false);
