run:
	go run main.go
swag:
	swag init docs/docs.go
compose:
	docker-compose up --build