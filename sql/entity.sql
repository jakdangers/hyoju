CREATE TABLE users
(
    id           BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL,
    deleted_at   TIMESTAMP NULL,
    nick_name    varchar(255),
    email        varchar(255),
    provider     varchar(255),
    firebase_uid varchar(255)
)