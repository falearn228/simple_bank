postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
startdb:
	docker start postgres17
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bobbabank
dropdb:
	docker exec -it postgres17 dropdb bobbabank	
migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/bobbabank?sslmode=disable" -verbose up
migrateup1:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/bobbabank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/bobbabank?sslmode=disable" -verbose down
migratedown1:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/bobbabank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockDB -destination internal/db/mock/store.go bobbabank/internal/db/sqlc Store
.PHONY: postgres start createdb dropdb migrateup migratedown sqlc test server mock