-- +goose Up

CREATE TABLE transaction
(
    id              UUID PRIMARY KEY    NOT NULL,
    amount          NUMERIC(14,0)       NOT NULL DEFAULT 0,  
    type            VARCHAR(20)         NOT NULL,
    unit            VARCHAR(20)         NOT NULL DEFAULT 'INR',
    status          VARCHAR(20)         NOT NULL,
    source_id       UUID                NOT NULL,
    destination_id  UUID                NOT NULL,
    comment         VARCHAR(225)                 DEFAULT NULL,
    created_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    updated_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    created_by      VARCHAR(100)        NOT NULL,
    updated_by      VARCHAR(100)        NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS transaction;