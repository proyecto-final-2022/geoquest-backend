package main

import (
	"geoquest-backend/cmd/api/routes"

	_ "geoquest-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title GeoQuest Backend
//@version 1.0
//@description App designed to communicate with GeoQuest's mobile app and provide CRUD functions
func main() {

	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
