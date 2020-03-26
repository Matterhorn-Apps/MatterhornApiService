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