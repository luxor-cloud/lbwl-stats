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
	map TINYTEXT NOT NULL,
	record_time BIGINT UNSIGNED NOT NULL,
	accomplished_at DATETIME NOT NULL,
	PRIMARY KEY (uuid, map),
	CONSTRAINT FK_map_records_global_game_data FOREIGN KEY (uuid) REFERENCES global_game_data (uuid)
)
COLLATE=utf8_general_ci
ENGINE=InnoDB;

CREATE TABLE checkpoint_records (
	uuid VARCHAR(36) NOT NULL,
	map TINYTEXT NOT NULL,
	checkpoint TINYINT UNSIGNED NOT NULL,
	record_time BIGINT UNSIGNED NOT NULL,
	accomplished_at DATETIME NOT NULL,
	INDEX i (map, checkpoint, uuid),
    CONSTRAINT FK_checkpoint_records_global_game_data FOREIGN KEY (uuid) REFERENCES global_game_data (uuid),
	CONSTRAINT FK_checkpoint_records_map_records FOREIGN KEY (map) REFERENCglobal_game_dataES map_records (map)
)
COLLATE=utf8_general_ci
ENGINE=InnoDB;
