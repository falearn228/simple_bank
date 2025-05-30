postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
startdb:
	docker start postgres17
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres17 dropdb simple_bank	
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockDB -destination db/mock/store.go github.com/simple_bank_course/db/sqlc Store
.PHONY: postgres start createdb dropdb migrateup migratedown sqlc test server mock