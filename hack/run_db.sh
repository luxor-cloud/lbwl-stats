#!/usr/bin/env bash
docker run --rm --name statspg \
  -d \
  -p 5432:5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=secret \
  -v $(pwd)/create_tables.sql:/docker-entrypoint-initdb.d/create_tables.sql \
  -v $(pwd)/populate.sql:/docker-entrypoint-initdb.d/populate.sql \
  postgres
