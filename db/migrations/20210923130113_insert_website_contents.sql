
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO website_contents (website_id, title, contents) VALUES (1, '', '金融に関する情報を発信するメディアサイト。金融リテラシーを身につけるために大きな壁となる「情報の非対称性」をなくし、全世界90億人が夢にチャレンジできる世界を実現することを目標とし、ユーザーに有益な情報を届けています。');
INSERT INTO website_contents (website_id, title, contents) VALUES (2, '', '日本の企業のデジタル情報格差問題を解決するべく、Webマーケティングを中心に情報を発信しているメディア。');
INSERT INTO website_contents (website_id, title, contents) VALUES (3, '', '「価値あるコンテンツ」をポリシーとした株式会社ルーシーが運営するwebメディア。どこよりも結果が出るマーケティングノウハウを発信している。');
INSERT INTO website_contents (website_id, title, contents) VALUES (4, '', 'ファッションを中心にビューティー、アート、グルメ、音楽、映画など、ライフスタイルに関する情報を発信している。');
INSERT INTO website_contents (website_id, title, contents) VALUES (5, '', '2018年にフォーブスの「Asia\’s 200 Best Under A Billion」に、2017年にFast Companyの「Most Innovative Companies」に選ばれたHypebeastが運営するwebメディア。ファッション、アート、デザイン、音楽といった領域で情報を発信する。');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from website_contents;
