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