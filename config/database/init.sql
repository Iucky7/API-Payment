CREATE DATABASE api_payment

CREATE TABLE merchant (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE payment (
    id VARCHAR(100) PRIMARY KEY,
    merchant_id VARCHAR(100) NOT NULL,
    bank_account VARCHAR(100) NOT NULL,
    amount bigint NOT NULL,
    payment_date timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY(merchant_id) REFERENCES merchant(id)
);

CREATE TABLE user_credential (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL
);