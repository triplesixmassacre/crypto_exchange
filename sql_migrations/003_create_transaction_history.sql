CREATE TABLE transaction_history (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    asset TEXT NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    trans_type TEXT NOT NULL, -- тип транзакции(деп/вывод/перевод)
    status TEXT NOT NULL, -- статус транзакции (confirmed/failed/penging)
    trans_hash TEXT, -- хэш транзакции в блокчейне(это дикпик сказал я не знаю что это)
    from_user_id BIGINT, -- from и to для переводов 
    to_user_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
); 