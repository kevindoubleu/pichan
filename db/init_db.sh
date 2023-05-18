#!/bin/sh

echo "initializing db contents"

setup_db() {
  cd scripts/$1

  psql -U $POSTGRES_USER -af create-db.sql
  psql -U $POSTGRES_USER -af create-types.sql
  psql -U $POSTGRES_USER -af create-tables.sql
  echo "dollar 2 is" $2
  if [ "$2" == "--sample" ]; then
    psql -U $POSTGRES_USER -af insert-data.sql
  fi

  cd ../..
}

main() {
  setup_db books $1
  setup_db habits $1
}

main "$@"
