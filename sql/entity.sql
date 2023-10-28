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
    code  varchar(255)
);

CREATE TABLE challenges
(
    id         INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    user_id  BINARY(16),
    title      VARCHAR(255),
    emoji      VARCHAR(255),
    duration   VARCHAR(255),
    start_date TIMESTAMP,
    end_date   TIMESTAMP,
    plan_time  TIMESTAMP,
    alarm      BOOLEAN,
    week_day   TINYINT UNSIGNED,
    type       VARCHAR(255),
    status     VARCHAR(255),
    code       VARCHAR(255)
);

CREATE TABLE challenge_participants
(
    id         INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    challenge_id INT UNSIGNED,
    user_id    BINARY(16)
)

CREATE TABLE challenge_histories
(
    id          INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP NULL,
    user_id     BINARY(16),
    challenge_id  INT UNSIGNED,
    plan_time   TIMESTAMP,
    front_image VARCHAR(255),
    back_image  VARCHAR(255)
)

