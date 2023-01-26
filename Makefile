postgres:
	docker compose up
	# docker run simple-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker compose exec -it simple-bank createdb --username=root --owner=root simple-bank
dropdb:
	docker compose exec -it simple-bank dropdb simple-bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple-bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://poastgres:password@localhost:5432/simple-bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc