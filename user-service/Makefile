DB_URL = postgres://postgres:1@localhost:5432/userdb?sslmode=disable

migrate-up:
	migrate -path ./db/migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./db/migrations -database ${DB_URL} down

migrate-force:
<<<<<<< HEAD
	migrate -path ./db/migrations -database ${DB_URL} force 1
=======
	migrate -path ./db/migrations -database ${DB_URL} force 1
>>>>>>> origin/main
