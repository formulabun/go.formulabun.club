CREATE TABLE IF NOT EXISTS replays (
  ReplayID int auto_increment primary key,
  GameMap INT2 unsigned,
  Time INT4 unsigned,
  BestLap INT4 unsigned,
  
  PlayerName CHAR(16),
  PlayerSkin CHAR(16),
  PlayerColor CHAR(16),
  Speed TINYINT UNSIGNED,
  Weight TINYINT UNSIGNED,
  UNIQUE (Time, BestLap, PlayerName)
);

CREATE TABLE IF NOT EXISTS replayfiles (
  ReplayID INT REFERENCES replays(ReplayID),
  FileName VARCHAR(512),
  FileCheckSum char(32),
  UNIQUE (ReplayID, FileName)
);
