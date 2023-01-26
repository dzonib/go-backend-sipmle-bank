postgres:
root username: root,
password: secret

docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# get shell for root user

docker exec -it postgres12 psql -U root

# quit shell

\q

# get logs from container

docker logs postgres12

# migration command (creates initial migrations up and down)

migrate create -ext sql -dir db/migration -seq init_schema

* `-ext sql` - extension sql 
* `-seq` generate sequential version number for the migration file

# execute shell when container is running

docker exec -it postgres12 /bin/sh

# create db

createdb --username=root --owner=root simple_bank

# switch to that db terminal

psql simple_bank

# remove db

dropdb simple_bank

# exit the psql shell

exit

# create and drop db without interacting with psql shell

docker exec -it postgres12 createdb --username=root --owner=root simple_bank

docker exec -it postgres12 dropdb simple_bank

# sqlc

sqlc init -> inits yaml file
