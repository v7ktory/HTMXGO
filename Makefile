.PHONY:
.SILENT:
.DEFAULT_GOAL := run
run:
	docker-compose up -d --remove-orphans webapp
stop: 
	docker-compose down -v
up:
	migrate -path ./migration -database 'postgresql://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
down:
	migrate -path ./migration -database 'postgresql://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
