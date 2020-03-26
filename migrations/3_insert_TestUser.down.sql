DELETE FROM Users WHERE UserID=1;

ALTER TABLE FoodRecords
    DROP FOREIGN KEY fk_FoodRecords_UserID,
    ADD FOREIGN KEY (UserID) REFERENCES Users(UserID);

ALTER TABLE ExerciseRecords
    DROP FOREIGN KEY fk_ExerciseRecords_UserID,
    ADD FOREIGN KEY (UserID) REFERENCES Users(UserID);

ALTER TABLE CalorieGoals
    DROP FOREIGN KEY fk_CalorieGoals_UserID,
    ADD FOREIGN KEY (UserID) REFERENCES Users(UserID);