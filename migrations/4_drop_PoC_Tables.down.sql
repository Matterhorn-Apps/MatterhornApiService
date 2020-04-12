CREATE TABLE IF NOT EXISTS Counters (
  ID int,
  Value int,
  PRIMARY KEY (ID)
);

INSERT IGNORE INTO Counters(ID, Value)
VALUES(1, 0);

CREATE TABLE IF NOT EXISTS Users (
  UserID int,
  DisplayName VARCHAR(50),
  Age int,
  Height int,
  Sex ENUM('Female', 'Male'),
  Weight int,
  PRIMARY KEY (UserID)
);

CREATE TABLE IF NOT EXISTS CalorieGoals (
    ID int AUTO_INCREMENT,
    UserID int NOT NULL,
    Calories int NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE IF NOT EXISTS ExerciseRecords (
    ID int AUTO_INCREMENT,
    UserID int NOT NULL,
    Calories int NOT NULL,
    Label VARCHAR(255) DEFAULT '',
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (ID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE IF NOT EXISTS FoodRecords (
    ID int AUTO_INCREMENT,
    UserID int NOT NULL,
    Calories int NOT NULL,
    Label VARCHAR(255) DEFAULT '',
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (ID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

ALTER TABLE CalorieGoals 
    DROP FOREIGN KEY CalorieGoals_ibfk_1,
    ADD CONSTRAINT fk_CalorieGoals_UserID FOREIGN KEY (UserID)
        REFERENCES Users(UserID)
        ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ExerciseRecords 
    DROP FOREIGN KEY ExerciseRecords_ibfk_1,
    ADD CONSTRAINT fk_ExerciseRecords_UserID FOREIGN KEY (UserID)
        REFERENCES Users(UserID)
        ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE FoodRecords 
    DROP FOREIGN KEY FoodRecords_ibfk_1,
    ADD CONSTRAINT fk_FoodRecords_UserID FOREIGN KEY (UserID)
        REFERENCES Users(UserID)
        ON UPDATE CASCADE ON DELETE CASCADE;

INSERT INTO Users (UserID, DisplayName, Age, Height, Weight, Sex) 
VALUES (1, "Test User", 23, 68, 160.5, 'Male')
ON DUPLICATE KEY UPDATE DisplayName='Test User', Age=23, Height=68, Weight=160.5, Sex='Male';