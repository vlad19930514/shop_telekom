postgres:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it some-postgres createdb --username=root --owner=root telekom

dropdb:
	docker exec -it some-postgres  dropdb telekom

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/telekom?sslmode=disable" -verbose down
 
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/telekom?sslmode=disable" -verbose up

sqlc:
	sqlc generate
