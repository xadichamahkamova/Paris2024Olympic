DB_URL = postgres://postgres:1@localhost:5432/medaldb?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/medal.proto

migrate-create:
	@migrate create -ext sql -dir ./db/migrations -seq medals_table

migrate-up:
	migrate -path ./db/migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./db/migrations -database ${DB_URL} down

migrate-force:
	migrate -path=db/migrations -database ${DB_URL} verbose -force 1

run:
	@go run cmd/main.go