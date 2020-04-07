CREATE DATABASE test;
USE test;

CREATE TABLE player_stats
(
    uuid           VARCHAR(36)  NOT NULL,
    wins           INT UNSIGNED NOT NULL,
    deaths         INT UNSIGNED NOT NULL,
    checkpoints    INT UNSIGNED NOT NULL,
    games_played   INT UNSIGNED NOT NULL,
    instant_deaths INT UNSIGNED NOT NULL,
    points         INT UNSIGNED NOT NULL,
    PRIMARY KEY (uuid)
)
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

CREATE TABLE player_map_scores
(
    uuid            VARCHAR(36)     NOT NULL,
    map             VARCHAR(100)    NOT NULL,
    time_needed     BIGINT UNSIGNED NOT NULL,
    accomplished_at DATETIME        NOT NULL,
    INDEX duration (time_needed),
    PRIMARY KEY (uuid, map, accomplished_at),
    FOREIGN KEY (uuid) REFERENCES global_game_data (uuid)
)
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

CREATE TABLE player_checkpoint_score
(
    uuid            VARCHAR(36)      NOT NULL,
    map             VARCHAR(100)     NOT NULL,
    checkpoint      TINYINT UNSIGNED NOT NULL,
    time_needed     BIGINT UNSIGNED  NOT NULL,
    accomplished_at DATETIME         NOT NULL,
    INDEX time_needed_per_checkpoint (time_needed, checkpoint),
    PRIMARY KEY (uuid, map, accomplished_at),
    FOREIGN KEY (uuid) REFERENCES player_map_scores (uuid)
)
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;
