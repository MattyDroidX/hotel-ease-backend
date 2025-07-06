#!/usr/bin/env bash
# -----------------------------------------------------------------
#  Hotel Ease – inicializador “um-clique” (Go + PostgreSQL)
# -----------------------------------------------------------------
#  Requisitos:
#    • PostgreSQL instalado localmente (serviço rodando)
#    • Go 1.22+ na PATH
#    • bash + psql na PATH (Windows: Git Bash funciona)
# -----------------------------------------------------------------
set -euo pipefail

# ---------- 0. Carrega .env (se existir) -------------------------
if [[ -f .env ]]; then
  echo "▶ Lendo .env"
  while IFS='=' read -r key value; do
    [[ $key =~ ^[[:space:]]*# || -z $key ]] && continue
    export "$key"="${value//$'\r'/}" 
  done < .env
fi

# ---------- 1. Valores padrão ------------------------------------
DB_URL=${DB_URL:-postgres://hotelease_user:admin@localhost:5432/hotelease?sslmode=disable}
PORT=${PORT:-8080}

# ---------- 2. Parse da URL --------------------------------------
regex='postgres://([^:]+):([^@]+)@([^:/]+):([0-9]+)/([^?]+)'
[[ $DB_URL =~ $regex ]] || { echo "❌ DB_URL inválido"; exit 1; }

DB_USER="${BASH_REMATCH[1]}"
DB_PASS="${BASH_REMATCH[2]}"
DB_HOST="${BASH_REMATCH[3]}"
DB_PORT="${BASH_REMATCH[4]}"
DB_NAME="${BASH_REMATCH[5]}"

# ---------- 3. Dependências --------------------------------------
for cmd in go psql; do
  command -v "$cmd" >/dev/null || { echo "❌ $cmd não encontrado na PATH"; exit 1; }
done

# ---------- 4. Testa se banco/usuário já existem -----------------
if PGPASSWORD="$DB_PASS" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c '\q' 2>/dev/null; then
  echo "✔️  Banco \"$DB_NAME\" e usuário \"$DB_USER\" já existem – pulando criação"
else
  echo "⚠️  Banco ou usuário ausentes – criando para você…"

  # superusuário para criação (padrão: postgres)
  SUPERUSER=${SUPERUSER:-postgres}
  SUPERPASS=${SUPERPASS:-}

  if [[ -z $SUPERPASS ]]; then
    # pede senha só uma vez, sem echo no terminal
    read -s -p "Senha do superusuário \"$SUPERUSER\" (deixe em branco se não tem): " SUPERPASS
    echo
  fi

  export PGPASSWORD="$SUPERPASS"

  # cria role se faltar
  psql -U "$SUPERUSER" -h "$DB_HOST" -p "$DB_PORT" -tc \
      "SELECT 1 FROM pg_roles WHERE rolname='$DB_USER'" |
  grep -q 1 || psql -U "$SUPERUSER" -h "$DB_HOST" -p "$DB_PORT" \
      -c "CREATE ROLE $DB_USER WITH LOGIN PASSWORD '$DB_PASS';" \
      || { echo "❌ Falha ao criar usuário. Verifique senha/privilégios."; exit 1; }

  # cria database se faltar
  psql -U "$SUPERUSER" -h "$DB_HOST" -p "$DB_PORT" -tc \
      "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" |
  grep -q 1 || psql -U "$SUPERUSER" -h "$DB_HOST" -p "$DB_PORT" \
      -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;" \
      || { echo "❌ Falha ao criar banco. Verifique senha/privilégios."; exit 1; }

  unset PGPASSWORD
  echo "✅ Banco e usuário criados com sucesso!"
fi

# ---------- 5. Exporta variáveis que o Go usa --------------------
export DB_URL PORT

# ---------- 6. Inicia a API --------------------------------------
echo "▶ API em http://localhost:$PORT  (Ctrl+C para sair)"
go run main.go &
PID=$!

sleep 2
URL="http://localhost:$PORT/swagger/index.html"
if command -v xdg-open >/dev/null; then xdg-open "$URL"
elif command -v open     >/dev/null; then open "$URL"; fi

wait $PID
