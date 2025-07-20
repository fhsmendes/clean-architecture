#!/bin/sh

set -e

LOG_FILE=/app/logs/app.log
MYSQL_HOST=mysql
MYSQL_PORT=3306

mkdir -p "$(dirname "$LOG_FILE")"
exec > "$LOG_FILE" 2>&1

# echo "[entrypoint] Aguardando MySQL em $MYSQL_HOST:$MYSQL_PORT..."
# until nc -z "$MYSQL_HOST" "$MYSQL_PORT"; do
#   echo "[entrypoint] Aguardando MySQL..."
#   sleep 2
# done

echo "[entrypoint] Baixando migrate com suporte a MySQL..."
mkdir -p /tmp/migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz \
  | tar -xz -C /tmp/migrate
chmod +x /tmp/migrate/migrate

echo "[entrypoint] Executando migrations..."
/tmp/migrate/migrate -path=sql/migrations -database "mysql://root:root@tcp($MYSQL_HOST:$MYSQL_PORT)/orders" -verbose up

echo "[entrypoint] Subindo aplicação com Wire..."
exec go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
