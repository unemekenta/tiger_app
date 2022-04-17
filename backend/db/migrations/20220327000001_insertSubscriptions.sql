
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE money_accounts ADD subscriptions_flg BOOLEAN NOT NULL DEFAULT false;

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    money_account_id INTEGER NOT NULL REFERENCES money_accounts(id) ON DELETE CASCADE,
    start_year INTEGER,
    start_month INTEGER,
    end_year INTEGER,
    end_month INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE subscriptions;
ALTER TABLE money_accounts DROP COLUMN subscriptions_flg;