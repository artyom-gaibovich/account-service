

LOCAL_BIN:=$(CURDIR)/bin

proto_install_deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

proto_get_deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

proto_generate:
	make proto_generate_account_api

proto_generate_account_api:
	mkdir -p pkg/account_v1
	protoc --proto_path api/account_v1 \
	--go_out=pkg/account_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/account_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/account_v1/account.proto






run:
	go run cmd/web/main.go

test:
	go test ./...

db_url ?= postgres://postgres:postgres@localhost:5432/account-service?sslmode=disable

install_golang_migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

generate_migration:
	migrate create -ext sql -dir db/migrations $(name)

db_migrate:
	migrate -database $(db_url) -path db/migrations up

db_rollback:
	migrate -database $(db_url) -path db/migrations down 1

db_force:
	migrate -database $(db_url) -path db/migrations force $(version)

db_test_migrate:
	migrate -database postgres://postgres:postgres@localhost:5433/blog_test?sslmode=disable  -path db/migrations up

db_test_rollback:
	migrate -database postgres://postgres:postgres@localhost:5433/blog_test?sslmode=disable  -path db/migrations down 1

db_test_force:
	migrate -database postgres://postgres:postgres@localhost:5433/blog_test?sslmode=disable -path db/migrations force $(version)

send:
	echo $(m)

.PHONY: test





