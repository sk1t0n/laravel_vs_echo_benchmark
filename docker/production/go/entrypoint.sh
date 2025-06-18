#!/bin/sh
set -e

goose up

for i in `seq 1 100`;
do
    PGPASSWORD=password psql -h pgbouncer -p 6432 -d postgres -U postgres \
        -c "INSERT INTO todos_go (title) VALUES ('Post $i');"
done

exec "$@"
