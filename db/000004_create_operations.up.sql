CREATE TABLE IF NOT EXISTS operations
(
    id              UUID PRIMARY KEY,
    from_account_id UUID              NOT NULL,
    to_account_id   UUID              NOT NULL,
    amount          BIGINT            NOT NULL,
    type            CHARACTER VARYING NOT NULL,
    created_at      TIMESTAMP         NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    created_by      UUID              NOT NULL,
    status          CHARACTER VARYING NOT NULL DEFAULT 'PENDING',

    FOREIGN KEY (from_account_id) REFERENCES accounts (id),
    FOREIGN KEY (to_account_id) REFERENCES accounts (id),
    FOREIGN KEY (created_by) REFERENCES users (id),
    CONSTRAINT amount_check CHECK (amount >= 0)
);