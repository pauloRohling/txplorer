CREATE TABLE IF NOT EXISTS accounts
(
    id         UUID PRIMARY KEY,
    balance    BIGINT            NOT NULL DEFAULT 0,
    user_id    UUID UNIQUE       NOT NULL,
    created_at TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status     CHARACTER VARYING NOT NULL DEFAULT 'ACTIVE',

    FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT balance_check CHECK (balance >= 0)
);