-- +goose Up

CREATE TABLE kyc_details
(
    id              UUID PRIMARY KEY    NOT NULL,
    govt_id_number  VARCHAR(225)        NOT NULL,
    expiry_date     TIMESTAMP                    DEFAULT NULL,  
    status          VARCHAR(20)         NOT NULL,
    verified_by     VARCHAR(100)        NOT NULL,
    created_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    updated_on      TIMESTAMP           NOT NULL DEFAULT Now(),
    created_by      VARCHAR(100)        NOT NULL,
    updated_by      VARCHAR(100)        NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS kyc_details;