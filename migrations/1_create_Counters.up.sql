CREATE TABLE IF NOT EXISTS Counters (
  ID int,
  Value int,
  PRIMARY KEY (ID)
);

INSERT IGNORE INTO Counters(ID, Value)
VALUES(1, 0);