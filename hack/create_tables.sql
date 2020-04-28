CREATE SCHEMA flash;

CREATE TABLE flash.player_stats
(
    uuid           UUID PRIMARY KEY NOT NULL,
    wins           INT  NOT NULL,
    deaths         INT  NOT NULL,
    checkpoints    INT  NOT NULL,
    games_played   INT  NOT NULL,
    instant_deaths INT  NOT NULL,
    points         INT  NOT NULL
);

CREATE TABLE flash.player_map_scores
(
    uuid            UUID                        NOT NULL,
    map             VARCHAR(100)                NOT NULL,
    time_needed     BIGINT                      NOT NULL,
    accomplished_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY (uuid, map, accomplished_at),
    FOREIGN KEY (uuid) REFERENCES flash.player_stats (uuid)
);

CREATE TABLE flash.player_checkpoint_scores
(
    uuid            UUID                        NOT NULL,
    map             VARCHAR(100)                NOT NULL,
    checkpoint      SMALLINT                    NOT NULL,
    time_needed     BIGINT                      NOT NULL,
    accomplished_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY (uuid, map, accomplished_at),
    FOREIGN KEY (uuid) REFERENCES flash.player_stats (uuid)
);

CREATE INDEX duration ON flash.player_map_scores (time_needed);
CREATE INDEX time_needed_per_checkpoint ON flash.player_checkpoint_scores (time_needed, checkpoint);
