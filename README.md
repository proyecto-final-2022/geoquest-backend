# geoquest-backend
Backend for GeoQuest App

Comandos instalacion de Swagger:
- go install github.com/swaggo/swag/cmd/swag@latest (para version de Go 1.18)
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

Comandos instalacion de Mongo:
- go get go.mongodb.org/mongo-driver

Comandos makefile:
 - make run: levantar la aplicacion 
 - make swag: guardar lo nuevo que hayas sumado en swagger 
 - make compose: levantar el docker compose

Link al swagger de la app luego de levantarla:
- http://localhost:8080/swagger/index.html 

Link a la documentacion de Gin Swagger
- https://github.com/swaggo/gin-swagger (docu swagger)

Link a la documentacion de GORM
- https://gorm.io/docs/