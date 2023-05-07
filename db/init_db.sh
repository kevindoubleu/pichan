#!/bin/sh

echo "initializing db contents"

setup_db() {
  cd $1
  psql -U $POSTGRES_USER -af create-db.sql
  psql -U $POSTGRES_USER -af create-types.sql
  psql -U $POSTGRES_USER -af create-tables.sql
  psql -U $POSTGRES_USER -af insert-data.sql
  cd ..
}

main() {
  setup_db books
  setup_db habits
}

main
