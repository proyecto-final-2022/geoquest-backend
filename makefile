run:
	go run main.go
swag:
	swag init docs/docs.go
compose:
	docker-compose up --build
install: 
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files
