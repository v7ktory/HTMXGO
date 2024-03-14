# Как запустить

1. Склонируйте репозиторий
2. Переименовать .env.example в .env
4. Откройте два терминала 
5. В первом терминале зайдите в папку ui и запустите команду `npx tailwindcss -i ./static/main.css -o ./static/tailwind.css --watch` (Если у вас нет `npx` то выполните команду `npm install -D tailwindcss`) 
6. Во втором терминале выполните команды `make` и `make up`
7. После выполненных действий выполните `go run ./cmd/app/main.go`
8. Перейти по http://localhost:3000/
9. `make down` чтобы откатить бд
   
