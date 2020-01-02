#!/usr/bin/env bash
docker run --name mar -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret mariadb:10.4
