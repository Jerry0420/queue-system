#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER $POSTGRES_BACKEND_USER WITH NOINHERIT PASSWORD '$POSTGRES_BACKEND_PASSWORD';
    CREATE USER $POSTGRES_MIGRATION_USER WITH NOINHERIT PASSWORD '$POSTGRES_MIGRATION_PASSWORD';
    CREATE USER $POSTGRES_VAULT_USER WITH NOINHERIT CREATEROLE SUPERUSER PASSWORD '$POSTGRES_VAULT_PASSWORD';
    CREATE DATABASE $POSTGRES_BACKEND_DB;

    REVOKE USAGE ON SCHEMA public FROM PUBLIC;

    GRANT USAGE ON SCHEMA public to $POSTGRES_MIGRATION_USER,$POSTGRES_BACKEND_USER;
    GRANT CREATE ON SCHEMA public to $POSTGRES_MIGRATION_USER;
    GRANT CREATE ON DATABASE $POSTGRES_BACKEND_DB to $POSTGRES_MIGRATION_USER;
EOSQL