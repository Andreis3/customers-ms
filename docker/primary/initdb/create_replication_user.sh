#!/bin/bash
set -e

# Cria o role de replicação se ainda não existir
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=postgres <<-EOSQL
    DO \$\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = '${REPLICATION_USER}') THEN
            CREATE ROLE ${REPLICATION_USER} WITH REPLICATION LOGIN PASSWORD '${REPLICATION_PASSWORD}';
        END IF;
    END
    \$\$;
EOSQL
