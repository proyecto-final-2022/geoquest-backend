# geoquest-backend
Backend for GeoQuest App

Comandos instalacion de Swagger:
- go install github.com/swaggo/swag/cmd/swag@latest (para version de Go 1.18)
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

Comandos makefile:
 - make run: levantar la aplicacion (go run main.go)
 - make swag: guardar lo nuevo que hayas sumado en swagger (swag init docs/docs.go)

Link a la docu de swagger (despues de levantar la app):
- http://localhost:8080/swagger/index.html 
- https://github.com/swaggo/gin-swagger (docu swagger)
