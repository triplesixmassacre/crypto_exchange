CREATE TABLE balances (
        user_id BIGINT,
        asset TEXT,  -- "BTC", "USDT"
        amount DOUBLE PRECISION,
        PRIMARY KEY (user_id, asset)
    );