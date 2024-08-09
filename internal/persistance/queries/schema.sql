CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts
(
    id      UUID PRIMARY KEY,
    balance BIGINT NOT NULL DEFAULT 0,

    CONSTRAINT balance_check CHECK (balance >= 0)
);

CREATE TABLE IF NOT EXISTS transactions
(
    id              UUID PRIMARY KEY,
    from_account_id UUID              NOT NULL,
    to_account_id   UUID              NOT NULL,
    amount          BIGINT            NOT NULL,
    timestamp       TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status          CHARACTER VARYING NOT NULL DEFAULT 'PENDING',

    FOREIGN KEY (from_account_id) REFERENCES accounts (id),
    FOREIGN KEY (to_account_id) REFERENCES accounts (id),
    CONSTRAINT amount_check CHECK (amount >= 0)
);