-- +goose Up
CREATE TABLE transaction_history (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    asset VARCHAR(10) NOT NULL,
    amount NUMERIC(20, 8) NOT NULL,
    trans_type VARCHAR(20) NOT NULL, -- deposit, withdrawal, transfer
    status VARCHAR(20) NOT NULL, -- confirmed, failed, pending
    trans_hash VARCHAR(255),
    from_user_id BIGINT REFERENCES users(id),
    to_user_id BIGINT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS transaction_history; 