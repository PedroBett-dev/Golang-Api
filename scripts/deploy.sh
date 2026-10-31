#!/usr/bin/env bash
set -euo pipefail

if [ $# -eq 0 ]; then
  echo "Uso: $0 \"descrição do commit\""
  exit 1
fi

description="$*"
timestamp="$(date '+%Y-%m-%d %H:%M:%S')"

cd "$(git rev-parse --show-toplevel)"

git add .

if git diff --cached --quiet; then
  echo "Nenhuma alteração para commitar."
  exit 0
fi

git commit -m "[$timestamp] $description"

echo "Commit criado: [$timestamp] $description"

git push

echo "Push realizado."
