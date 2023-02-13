CREATE SCHEMA IF NOT EXISTS clients;

CREATE TABLE IF NOT EXISTS clients.users
(
    id                 BIGSERIAL PRIMARY KEY,
    first_name         TEXT NOT NULL,
    last_name          TEXT NOT NULL,
    commission_fix     NUMERIC DEFAULT 0,
    commission_percent NUMERIC DEFAULT 0
);

CREATE TABLE IF NOT EXISTS clients.balance
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id   INTEGER REFERENCES clients.users (id),
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS clients.transactions
(
    id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id          INTEGER REFERENCES clients.users (id),
    type               TEXT    NOT NULL,
    amount             NUMERIC NOT NULL,
    commission_fix     NUMERIC,
    commission_percent NUMERIC
);

CREATE TABLE IF NOT EXISTS clients.ledger
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id    UUID REFERENCES clients.balance (id),
    base_amount   NUMERIC NOT NULL,
    action_amount NUMERIC NOT NULL,
    action_type   TEXT    NOT NULL
);