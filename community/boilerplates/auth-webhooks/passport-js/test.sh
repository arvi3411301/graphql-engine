#!/usr/bin/env bash

set -evo pipefail
export DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASS}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}
echo "${DATABASE_URL}"
echo "Migrate database"
knex migrate:latest
echo "Migrated database"
