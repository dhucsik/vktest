migrate:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir=${PWD}/migrations up

docs:
	swag init -d ./internal/transport/http -g server.go --pd --parseDepth 2 -o ./swagger/docs