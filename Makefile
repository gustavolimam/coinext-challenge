.DEFAULT_GOAL := run

postgres:
	@docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=gustavolimam -e POSTGRES_PASSWORD=gustavolimam -d postgres:12-alpine

# createdb:
#     docker exec -it postgres12 createdb --username=gustavolima zombieland

# dropdb:
#     docker exec -it postgres12 dropdb zombieland

# migrateup:
#     @migrate -path db/migrations -database 'postgres://postgres:postgres@localhost:5432/zombieland?sslmode=disable' -verbose up

# migratedown:
#     migrate -path db/migration -database 'postgresql://gustavolimam:gustavolimam@localhost:5432/simple_bank?sslmode=disable' -verbose down

run: 
	go mod tidy && go run cmd/main.go

mock_repository:
	cd internal/repository && mockery --all --output=mocks

mock_services:
	cd internal/services && mockery --all --output=mocks

test:
	@echo Running tests with coverate report...
	go test -v -coverprofile=$(COVER_OUT) $(shell go list ./...) -json 2>&1 | tee /tmp/gotest.log | gotestfmt -hide successful-tests,empty-packages