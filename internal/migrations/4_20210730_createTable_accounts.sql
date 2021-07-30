-- +goose Up

CREATE TABLE accounts
(
    id              UUID PRIMARY KEY    NOT NULL,
    balance         NUMERIC(14,0)       NOT NULL DEFAULT 0,  
    type            VARCHAR(20)         NOT NULL,
    unit            VARCHAR(20)         NOT NULL DEFAULT 'INR',
    created_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    updated_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    created_by      VARCHAR(100)        NOT NULL,
    updated_by      VARCHAR(100)        NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS accounts;