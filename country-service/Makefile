DB_URL = postgres://postgres:1@localhost:5432/countrydb?sslmode=disable

migrate-create:
	@migrate create -ext sql -dir ./db/migrations -seq countries_table

migrate-up:
	migrate -path ./db/migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./db/migrations -database ${DB_URL} down

migrate-force:
	migrate -path ./db/migrations -database ${DB_URL} force 1	