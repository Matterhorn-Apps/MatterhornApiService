ALTER TABLE users
DROP COLUMN weight_pounds;

CREATE TABLE weight_records (
    PRIMARY KEY (weight_record_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    user_id                 VARCHAR(255)                        NOT NULL,
    weight_record_id        INT UNSIGNED AUTO_INCREMENT         NOT NULL,
    weight_pounds           INT UNSIGNED                        NOT NULL,
    timestamp               TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);