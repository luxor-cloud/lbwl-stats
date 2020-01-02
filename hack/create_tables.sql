CREATE DATABASE test;
USE test;

CREATE TABLE global_game_data (
	uuid VARCHAR(36) NOT NULL,
	wins INT UNSIGNED NOT NULL,
	deaths INT UNSIGNED NOT NULL,
	checkpoints INT UNSIGNED NOT NULL,
	games_played INT UNSIGNED NOT NULL,
	instant_deaths INT UNSIGNED NOT NULL,
	PRIMARY KEY (uuid)
)
COLLATE=utf8_general_ci
ENGINE=InnoDB;

CREATE TABLE map_records (
	uuid VARCHAR(36) NOT NULL,
	map VARCHAR(100) NOT NULL,
	record_time BIGINT UNSIGNED NOT NULL,
	accomplished_at DATETIME NOT NULL,
	PRIMARY KEY (map, uuid),
	FOREIGN KEY (uuid) REFERENCES global_game_data (uuid)
)
COLLATE=utf8_general_ci
ENGINE=InnoDB;

CREATE TABLE checkpoint_records (
	uuid VARCHAR(36) NOT NULL,
	map VARCHAR(100) NOT NULL,
	checkpoint TINYINT UNSIGNED NOT NULL,
	record_time BIGINT UNSIGNED NOT NULL,
	accomplished_at DATETIME NOT NULL,
	PRIMARY KEY (map, checkpoint, uuid),
	FOREIGN KEY (uuid, map) REFERENCES map_records (uuid, map)
)
COLLATE=utf8_general_ci
ENGINE=InnoDB;
