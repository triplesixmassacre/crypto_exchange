-- +goose Up
CREATE TABLE balances (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    currency VARCHAR(10) NOT NULL,
    amount NUMERIC(20, 8) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, currency)
);

-- +goose Down
DROP TABLE IF EXISTS balances;
