# Как запустить

1. Склонируйте репозиторий
2. Откройте два терминала 
3. В первом терминале зайдите в папку ui и запустите команду `npx tailwindcss -i ./static/main.css -o ./static/tailwind.css --watch` (Если у вас нет `npx` то выполните команду `npm install -D tailwindcss`) 
4. Во втором терминале выполните команды `make` и `make up`
5. После выполненных действий выполните `go run ./cmd/app/main.go`
6. `make down` чтобы откатить бд
