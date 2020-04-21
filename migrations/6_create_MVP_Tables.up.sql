CREATE TABLE users (
    PRIMARY KEY (user_id),
    user_id         VARCHAR(255)            NOT NULL,
    age             INT UNSIGNED,
    calorie_goal    INT UNSIGNED,
    display_name    VARCHAR(255),
    height_inches   INT UNSIGNED,
    sex             ENUM('Female', 'Male', 'Other'),
    weight_pounds   INT UNSIGNED
);

CREATE TABLE exercise_records (
    PRIMARY KEY (exercise_record_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    user_id                 VARCHAR(255)                        NOT NULL,
    exercise_record_id      INT UNSIGNED AUTO_INCREMENT         NOT NULL,
    calories                INT UNSIGNED                        NOT NULL,
    label                   VARCHAR(255),
    timestamp               TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE food_records (
    PRIMARY KEY (food_record_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    user_id                 VARCHAR(255)                        NOT NULL,
    food_record_id          INT UNSIGNED AUTO_INCREMENT         NOT NULL,
    calories                INT UNSIGNED                        NOT NULL,
    label                   VARCHAR(255),
    timestamp               TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);