
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO money_account_labels (id, name) VALUES (1,'income');
INSERT INTO money_account_labels (id, name) VALUES (2,'expenses');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from money_account_labels WHERE id BETWEEN 1 AND 2;
SELECT setval('money_account_labels_id_seq', 1, false);
