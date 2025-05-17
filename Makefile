export

DB_MIGRATE_URL = postgres://login:pass@localhost:5432/postgres?sslmode=disable
MIGRATE_PATH = ./migration/postgres

up:
	docker compose up --build --force-recreate

down:
	docker compose down

test:
	go test -v -cover ./...

integration-test-http-v1:
	go test -count=1 -v -tags=integration,http_v1 ./test/integration

integration-test-http-v2:
	go test -count=1 -v -tags=integration,http_v2 ./test/integration

integration-test-grpc:
	go test -count=1 -v -tags=integration,grpc ./test/integration

integration-test: integration-test-http-v1 integration-test-http-v2 integration-test-grpc

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	@read -p "Name:" name; \
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $$name

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all

oapi-install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

generate: grpc_gen
	go generate ./...

grpc_gen:
	mkdir -p ./gen/grpc/profile_v1
	./bin/protoc \
	  --proto_path=api/grpc \
	  --go_out=./gen/grpc/profile_v1 --go_opt=paths=source_relative \
	  --plugin=protoc-gen-go=./bin/protoc-gen-go \
	  --go-grpc_out=./gen/grpc/profile_v1 --go-grpc_opt=paths=source_relative \
	  --plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	  profile_v1.proto
