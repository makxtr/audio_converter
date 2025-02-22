up:
	docker-compose --env-file .env.local up -d

down:
	docker-compose --env-file .env.local down --remove-orphans

build:
	docker-compose --env-file .env.local build --no-cache

restart:
	docker-compose --env-file .env.local restart

logs:
	docker-compose --env-file .env.local logs -f

test:
	docker-compose --env-file .env.local run --rm --entrypoint go go-app test ./... -v

lint:
	docker-compose run --rm go-app golangci-lint run ./...

format:
	docker-compose run --rm go-app gofmt -w .

run:
	docker-compose run --rm go-app go run main.go
