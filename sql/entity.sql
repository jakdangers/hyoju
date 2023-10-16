CREATE TABLE users
(
    id           BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL,
    deleted_at   TIMESTAMP NULL,
    nick_name    varchar(255),
    email        varchar(255),
    provider     varchar(255),
    firebase_uid varchar(255),
    friend_code  varchar(255)
)

CREATE TABLE Mission
(
    id         BIGINT PRIMARY KEY,
    author_id  BINARY(16),
    title      VARCHAR(255),
    emoji      VARCHAR(255),
    duration   VARCHAR(255),
    start_date DATETIME,
    end_date   DATETIME,
    plan_time  DATETIME,
    alarm      BOOLEAN,
    days       TINYINT UNSIGNED
);
