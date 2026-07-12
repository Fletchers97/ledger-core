CREATE TABLE accounts (
    id VARCHAR(255) PRIMARY KEY,
    balance BIGINT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);