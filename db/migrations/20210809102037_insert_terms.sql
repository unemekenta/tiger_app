
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO terms (name, content) VALUES ('サーバー', '利用者の要求（リクエスト）に対して、それに応答したデータを提供するコンピュータやプログラムのことを“サーバー”と呼んでいます。');
INSERT INTO terms (name, content) VALUES ('SSL', '通信プロトコルの一種で、第三者からの覗き見や書き換えを防止したりするためのもの。');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from terms WHERE name = 'サーバー';
DELETE from terms WHERE name = 'SSL';