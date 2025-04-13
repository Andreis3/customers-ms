#!/bin/bash
set -e

# Se a pasta de dados já existir, não faz nada.
if [ "$(ls -A "$PGDATA")" ]; then
  echo "Dados já existem em \$PGDATA, iniciando o PostgreSQL..."
  exit 0
fi

echo "Realizando pg_basebackup do primário..."

# Realiza o clone usando o pg_basebackup
pg_basebackup -h ${PRIMARY_HOST} -D "$PGDATA" -U ${REPLICATION_USER} -v -P --wal-method=stream

# Cria o arquivo standby.signal para indicar que é uma réplica (PostgreSQL 12+)
touch "$PGDATA/standby.signal"

# Cria o arquivo de configuração para conexão com o primário
cat > "$PGDATA/postgresql.auto.conf" <<EOF
primary_conninfo = 'host=${PRIMARY_HOST} port=5432 user=${REPLICATION_USER} password=${REPLICATION_PASSWORD}'
EOF
