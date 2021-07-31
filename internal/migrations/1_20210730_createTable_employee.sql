-- +goose Up

CREATE TABLE employee
(
    id          UUID PRIMARY KEY    NOT NULL,
    name        VARCHAR(225)        NOT NULL,
    email       VARCHAR(100)        NOT NULL,
    password	VARCHAR(225)        NOT null,
    type        VARCHAR(20)         NOT NULL,
    created_on  TIMESTAMP           NOT NULL DEFAULT Now(),
    updated_on  TIMESTAMP           NOT NULL DEFAULT Now(),
    created_by  VARCHAR(100)        NOT NULL,
    updated_by  VARCHAR(100)        NOT NULL
);

CREATE UNIQUE INDEX employee_idx_email ON employee(email);

-- +goose Down
DROP TABLE IF EXISTS employee;