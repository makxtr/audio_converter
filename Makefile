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

DB_URL=mysql://root:rootpass@db_mysql_user:3306/project_db

migrate-diff:
	docker-compose --env-file .env.local exec go-app sh -c 'atlas migrate diff new_changes --to file://db/schema.hcl --dev-url "mysql://root:rootpass@db_mysql_user:3306/project_db?parseTime=true"'


migrate-up:
	docker-compose exec go-app goose -dir db/migrations mysql "$(DB_URL)" up

migrate-down:
	docker-compose exec go-app goose -dir db/migrations mysql "$(DB_URL)" down

migrate-new:
	@mkdir -p db/migrations
	@read -p "Enter migration name: " name; \
	docker-compose exec go-app goose -dir db/migrations create "$$name" sql