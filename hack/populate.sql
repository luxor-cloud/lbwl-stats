
------------------------------------------------
-- player_stats
------------------------------------------------

INSERT INTO flash.player_stats (uuid, wins, deaths, checkpoints, games_played, instant_deaths, points)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 1337, 12315252, 245245, 5245245, 3534535, 15);

INSERT INTO flash.player_stats (uuid, wins, deaths, checkpoints, games_played, instant_deaths, points)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 4314, 2314, 245245, 134134, 531515, 135);


------------------------------------------------
-- player_map_scores
------------------------------------------------

-- MAP 1
-- 92de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 0, '2025-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, '2014-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 3, '2016-12-10 14:30:23.145890');

-- MAP 2
-- 92de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 0, '2028-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, '2020-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 3, '2018-12-10 14:30:23.145890');


-- MAP 1
-- 82de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, '2215-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, '2514-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 3, '2116-12-10 14:30:23.145890');


-- MAP 1
-- 82de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, '2015-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, '2014-12-10 14:30:23.145890');

INSERT INTO flash.player_map_scores (uuid, map, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 3, '2016-12-10 14:30:23.145890');


------------------------------------------------------------
-- player_checkpoint_scores
------------------------------------------------------------

-- MAP 1 --  multiple checkpoint scores
-- 92de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 1, '2017-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 2, '2016-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 3, '2018-03-05 14:00:23.145890');


INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 1, '2017-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 2, '2016-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 3, '2018-03-05 15:00:23.145890');

-- MAP 2 --  multiple checkpoint scores
-- 92de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 1, '2020-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 2, '2024-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 3, '2016-03-05 14:00:23.145890');


INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 1, '2018-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 2, '2020-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('92de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 3, '2015-03-05 15:00:23.145890');


-- MAP 1 --  multiple checkpoint scores
-- 82de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 1, '2020-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 2, '2024-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 1, 3, '2016-03-05 14:00:23.145890');


INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 2000, '2018-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 500, '2020-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map1', 2, 20, '2015-03-05 15:00:23.145890');

-- MAP 2 --  multiple checkpoint scores
-- 82de217b-8b2b-403b-86a5-fe26fa3a9b5f

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 1, '2026-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 1, '2025-03-05 14:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 1, 3, '2017-03-05 14:00:23.145890');


INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 1, '2038-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 2, '2010-03-05 15:00:23.145890');

INSERT INTO flash.player_checkpoint_scores (uuid, map, checkpoint, time_needed, accomplished_at)
VALUES ('82de217b-8b2b-403b-86a5-fe26fa3a9b5f', 'Map2', 2, 3, '2085-03-05 15:00:23.145890');
