-- migrate:up

CREATE TABLE cars
(
    id               SERIAL       NOT NULL UNIQUE,
    reg_num          CHAR(9)      NOT NULL UNIQUE,
    mark             VARCHAR(255) NOT NULL,
    model            VARCHAR(255) NOT NULL,
    year             SMALLINT,
    owner_name       VARCHAR(255) NOT NULL,
    owner_surname    VARCHAR(255) NOT NULL,
    owner_patronymic VARCHAR(255)

);

-- migrate:down

DROP TABLE cars;