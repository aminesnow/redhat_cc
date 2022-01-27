#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" \
    -v bucket_store_admin_pwd="$BUCKET_STORE_ADMIN_PWD" -f /sql/db.sql

psql -v ON_ERROR_STOP=1 --username "$BUCKET_STORE_ADMIN_USR" --dbname "$BUCKET_STORE_DB" -f /sql/tables.sql