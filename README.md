# Как запустить

1. Склонируйте репозиторий
2. Добавить .env {
3. API_KEY="5YxDmd6E2STUZoIZyCfLHg==0u2DvDIAeRGwb53f"

POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="qwerty"
POSTGRES_DB_NAME="postgres"
POSTGRES_SSL_MODE="disable"

REDIS_PASSWORD="qwerty"
REDIS_ADDR="localhost:6379"
REDIS_DB=0
}
4. Откройте два терминала 
5. В первом терминале зайдите в папку ui и запустите команду `npx tailwindcss -i ./static/main.css -o ./static/tailwind.css --watch` (Если у вас нет `npx` то выполните команду `npm install -D tailwindcss`) 
6. Во втором терминале выполните команды `make` и `make up`
7. После выполненных действий выполните `go run ./cmd/app/main.go`
8. Перейти по http://localhost:3000/
9. `make down` чтобы откатить бд
   
