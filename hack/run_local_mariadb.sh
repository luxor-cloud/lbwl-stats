#!/usr/bin/env bash
docker run --name mar \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=secret \
  -v $(pwd)/create_tables.sql:/docker-entrypoint-initdb.d/create_tables.sql \
  -v $(pwd)/populate_mariadb.sql:/docker-entrypoint-initdb.d/populate_mariadb.sql \
  mariadb:10
