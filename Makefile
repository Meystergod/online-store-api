APP_BIN = application/build/application

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./application/cmd/application/main.go

.PHONY: clean
clean:
	rm -rf ./application/build || true

.PHONY: swagger
swagger:
	swag init -g ./application/cmd/application/main.go -o ./application/docs

.PHONY: migrate
migrate:
	$(APP_BIN) migrate -version $(version)

.PHONY: migrate.down
migrate.down:
	$(APP_BIN) migrate -seq down

.PHONY: migrate.up
migrate.up:
	$(APP_BIN) migrate -seq up
