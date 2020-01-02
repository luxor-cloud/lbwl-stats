#!/usr/bin/env bash
docker exec -i mar sh -c 'exec mysql -uroot -p"secret"' < create_tables.sql
docker exec -i mar sh -c 'exec mysql -uroot -p"secret"' < populate_mariadb.sql