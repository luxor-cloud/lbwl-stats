USE test;

/*
INSERT INTO global_game_data (uuid, wins, deaths, checkpoints, games_played, instant_deaths)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 1337, 12315252, 245245, 5245245, 3534535);

INSERT INTO map_records (uuid, map, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 22324234, '2019-01-01 02:59:23.145890');

INSERT INTO map_records (uuid, map, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2948471, '2014-12-10 14:30:23.145890');

INSERT INTO map_records (uuid, map, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map3', 40304034, '2018-03-05 13:59:23.145890');
*/



INSERT INTO player_stats (uuid, wins, deaths, checkpoints, games_played, instant_deaths, points)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 1337, 12315252, 245245, 5245245, 3534535, 15);

INSERT INTO player_stats (uuid, wins, deaths, checkpoints, games_played, instant_deaths, points)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 1337, 12315252, 245245, 5245245, 3534535, 15);


-- Map1
INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 0, '2025-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, '2014-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 3, '2016-12-10 14:30:23.145890');


INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, '2215-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, '2514-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 3, '2116-12-10 14:30:23.145890');




-- Map2
INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, '2215-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, '2214-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 3, '2116-12-10 14:30:23.145890');


INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, '2015-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, '2014-12-10 14:30:23.145890');

INSERT INTO player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 3, '2016-12-10 14:30:23.145890');



/*
-- CHECKPOINT 1
INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 1, '2017-03-05 14:00:23.145890');

INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 2, '2016-03-05 14:00:23.145890');

INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 3, '2018-03-05 14:00:23.145890');

-- CHECKPOINT 2
INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 1, '2017-03-05 15:00:23.145890');

INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 2, '2016-03-05 15:00:23.145890');

INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 3, '2018-03-05 15:00:23.145890');
*/



/*
INSERT INTO player_checkpoint_score (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 3, 53424, '2018-03-05 16:00:23.145890');
*/

/*
INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 83424, '2018-03-05 17:00:23.145890');

INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 93424, '2018-03-05 18:00:23.145890');

INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 3, 123424, '2018-03-05 19:00:23.145890');


INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map3', 1, 143424, '2018-03-05 20:00:23.145890');

INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map3', 2, 183424, '2018-03-05 21:00:23.145890');

INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map3', 3, 253424, '2018-03-05 22:00:23.145890');


INSERT INTO global_game_data (uuid, wins, deaths, checkpoints, games_played, instant_deaths)
VALUES ('aabb023c-66ef-4168-9bf9-b898bee9f73f', 1, 20, 40, 100, 50);

INSERT INTO map_records (uuid, map, record_time, accomplished_at)
VALUES ('aabb023c-66ef-4168-9bf9-b898bee9f73f', 'Map1', 1337, '2015-01-01 02:59:23.145890');

INSERT INTO checkpoint_records (uuid, map, checkpoint, record_time, accomplished_at)
VALUES ('aabb023c-66ef-4168-9bf9-b898bee9f73f', 'Map1', 1, 25565, '2014-03-05 14:00:23.145890');
*/
