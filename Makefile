APP_BIN = application/build/application

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./application/cmd/application/main.go

clean:
	rm -rf ./application/build || true

swagger:
	swag init -g ./application/cmd/application/main.go -o ./application/docs

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up
