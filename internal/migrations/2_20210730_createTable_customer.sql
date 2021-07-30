-- +goose Up

CREATE TABLE customer
(
    id              UUID PRIMARY KEY    NOT NULL,
    name            VARCHAR(225)        NOT NULL,
    kyc_detials_id  UUID                         DEFAULT NULL,  
    account_detials JSON                         DEFAULT '{}'::JSONB,    
    created_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    updated_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    created_by      VARCHAR(100)        NOT NULL,
    updated_by      VARCHAR(100)        NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS customer;