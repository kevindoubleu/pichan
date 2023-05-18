# pichan

Pi Central Home Assistant Network

## Running

Make sure you have:
- initialized the config files with `make init-configs` and modified the
  values accordingly if needed
- compiled the go binary with `make build-habits`

before running `docker compose up`

## DB

### Initializing

For the first time spinning up the docker containers, the db would be empty.
This is done on purpose to give the ability to initialize/reinitialize the database
whenever needed.

*Note*: All other containers depending on this DB will fail
upon the first run before the db is initialized

Initialize/reinitialize the db by going into the postgres container
(default name is `pichan-postgres-1`) with

```sh
docker exec -it pichan-postgres-1 sh
```

and running `/app/init_db.sh`. Running with `--sample` will also insert some sample data if any

After the db is initialized, restart the docker containers with `ctrl+c` and `docker compose up`

### Persistence

DB container's `/var/lib/postgresql/data` is mounted to the host's `db/data` in this repo.
As long as:
- that folder is not deleted
- the built postgres docker image is not deleted
- the db initialization script is not rerun

then the data will persist
